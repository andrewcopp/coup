# Machine Learning Engineer Nanodegree
## Capstone Proposal
Andrew Copp
September 8, 2017

## Proposal

### Domain Background
Game-playing agents are a popular subject in the the world of machine learning and artificial intelligence. [Deep Blue](http://ac.els-cdn.com/S0004370201001291/1-s2.0-S0004370201001291-main.pdf?_tid=a7a55a4a-94f9-11e7-92e4-00000aab0f26&acdnat=1504918743_9160cb43671014cf2918442085ba4dce) was one of the first to capture the public imagination. [AlphaGo](http://airesearch.com/wp-content/uploads/2016/01/deepmind-mastering-go.pdf) recently conquered the game of Go which was considered to be one of the last bastions for human superiority because of the massive state space. In the last month, the [OpenAI](https://blog.openai.com/dota-2/) project was able to beat the top Dota2 competitors in 1v1 games. This is a monumental achievement because, unlike Chess and Go, Dota2 is a game with hidden information. I would like to make a contribution to this space with a small project that play stochastic game with hidden information.

I am personally interested in completing a project that requires Reinforcement Learning. I've already use some Supervised Learning at my place of employment and I think I will start making use Unsupervised Learning and Deep Learning shortly. I am dual enrolled in the AIND program and, entering Term 2, am focusing on Deep Learning. Reinforcement Learning doesn't seem to be as pragmatic for the challenges I am facing at the moment. For that reason, I would like to get some more practice. Reinforcement Learning is a great tool for training game-playing agents and I happen to love board games. I want to use this opportunity to use Reinforcement Learning to build a game-playing agent for a lesser known game and possibly make me the first person to solve this particular problem.

The board game for this project will be Coup. A summary of the rules can be found in Appendix A.

### Problem Statement
The nature of Coup makes it a difficult to play for both humans and computers. Deciding when to bluff or when to challenge is a combination of reading your opponent and assessing the situation. While using computer vision to measure weaknesses in an opponent would be interesting, it will also be a great exercise to see if a learner can be trained to beat opponents simply by understanding situations more thoroughly.

### Datasets and Inputs
There are no public datasets for Coup. It appears the only dataset is the one use by the Coup mobile app to display stats and that is certainly private. Fortunately, reinforcement learning does not require prior data because we can generate all the data we need through simulations. It seems to be a point of pride amongst AI researchers to be able to build game-playing agents that start off completely randomly and learn from their early mistakes.

The game itself will only need a few inputs to start a simulation. Specifically, it will need to know which players are playing. Internally, the game will need to keep track of a lot of states. The state space of the game will be significantly larger than what was built for the Smartcab project.

_Sample Space Calculations_

The back-of-the-envelope calculation is a little daunting.

Players = 5 // Five players can play at once.
Card = 7 (Unknown, Duke, Assassin, Ambassador, Captain, Contessa, None) // Five types of cards plus hidden cards and inactive players.
Coins = 13 (0 - 12) // Player can have between zero and twelve coins.

Actions = 8 (Income + ForeignAid + Tax + Exchange + Assassinate + Steal + Coup + None)
Challenge = 20 (Players * Successful * NotChallenged)
Block = 40 (NotBlocked * Challenge)
Select = 14 (2 * Cards)

Simple State Space = (5 Players) * 2 * (7 Cards) * (13 Coins) * (8 Actions) * (20 Challenge) * (40 Block) * (2 Cards) = 81,536,000

Almost ~81.5 million states to search through! That's probably not that big for larger enterprise projects but that seems like a lot for me to handle on my first go. Let's see if we can get that down with some optimizations. After all, not all actions can be challenged or blocked.

Players = 5 // Five players can play at once.
Card = 7 (Unknown, Duke, Assassin, Ambassador, Captain, Contessa, None) // Five types of cards plus hidden cards and inactive players.
Coins = 13 (0 - 12) // Player can have between zero and twelve coins.

SuccessfulChallenge = 4 (Players - 1)
UnsuccessfulChallenge = 4 (Players - 1)
NoChallenge = 1
Challenge = 9 (SuccessfulChallenge + UnsuccessfulChallenge + NoChallenge)

Select = 12 (2 * (Cards - None))

Blocked = 9 (Challenge)
NotBlocked = 1
Block = 10

Income = 1
ForeignAid = 10 (Block)
Tax = 9 (Challenge)
Exchange = 64 (SuccessfulChallenge + Select(UnsuccessfulChallenge + NoChallenge))
Assassinate = 54 (SuccessfulChallenge * Block(UnsuccessfulChallenge + NoChallenge))
Steal = 54 (SuccessfulChallenge * Block(UnsuccessfulChallenge + NoChallenge))
Coup = 1
None = 1

Actions = 194 (Income + ForeignAid + Tax + Exchange + Assassinate + Steal + Coup + None)
Moves = 1940 (2 * Players * Actions)

State Space = Players * 2 * Card * Coins * Moves = 1,765,400

Just under two million states. That's almost a 50x reduction of the state space.

### Solution Statement
The problem will be solved by using Q learning to determine the optimal policy. An off-policy method will be the way to approach this particular problem because there are so many possible ways to approach a situation. It seems way more reasonable to make no potentially incorrect assumptions about the optimal policy and instead allow a learner to explore the state space. The value-function approximator will be the win percentage of games from previous rounds. This Monte Carlo tree search implementation was chosen because the lack of historical data about the game would make it difficult to train a neural net or supervised learner about situations.

### Benchmark Model
There seems to be only one game playing agent for Coup and that is the AI for the official Coup mobile app. The AI in this app seems to behave at random and plays without any regard for the current situation or the history of moves. A better solution is definitely possible. Because the rules of the Coup mobile app are slightly different than the actual game, a random opponent will be built to serve as the benchmark.

### Evaluation Metrics
There are five metrics that will be used to evaluate success.
1. Win Percentage - What percentage of the time does the agent win against a random opponent? A human opponent?
2. Bluff Precision - What percentage of the opponent's challenges are successful?
3. Bluff Recall - What percentage of bluffs made by the agents are caught by the opponent?
4. Challenge Precision - What percentage of the agent's challenges are successful?
5. Challenge Recall - What percentage of bluffs made by the opponent are caught by the agent?

### Project Design
The project can be broken down into three stages.

Three Stages
1. Setup
2. Learning
3. Testing

_Setup_
The beginning of this project will be codifying the rules of Coup into a playable game. While we want the learner to figure the game out on its own, we need to make sure there are hard and fast rules so that the learner can figure out where the edge of the playing field is.

_Learning_
Once the game has all of the rules established, we can set the learner loose. The original learner will assume equal probabilities of winning from all moves. At the end of each game, the probabilities will be updated to reflect what states have led to victories. Early on in the learning, states will be selected at random even as dominant strategies emerge. As the learning rate decays over time, the stronger moves will be utilized again and again.

The agents value-function approximator will be saved at the end of each round. There will be a historical record of every iteration of the agent which means we can use experience replay to train the next version of the agent.

_Testing_
At the conclusion of learning, the agent will be ready for testing. It will play a set number of games against previous iterations of itself and record the results. It will also be spot checked by the development team and board game enthusiasts.

### Appendix A

Coup is a game for two to five players. The object of the game is to eliminate other players and be the last player standing. At the beginning of the gam, each player is dealt two hidden role cards which represent lives. On their turn, a player announces which of six actions they choose to take. Some actions are limited to certain role cards but, since role cards are hidden, a player may choose to take the risk of bluffing a card they don't have in seeking a greater reward. Bluffing exposes a player to being challenged. When challenged, a player must present the card needed to perform their action. If a player is able to produce the necessary card, the challenge loses on of their cards. If the player cannot produce the necessary card, their action is blocked and they most forfeit a card. Cards that have been discarded are turned face up so as to publicly display information about what cards are no longer in circulation.
