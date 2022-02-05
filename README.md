#GenPC

Discord bot for automating Numenera (and eventually other system) tables.

Still in early development, for now commands are as follows:

Numenera Helper Commands

!info ARG: general overview of requested feature
ARG values: bot, cypher, oddity, artifact

!gen ARG: generates a random instance of the requested feature
ARG values: oddity, (soon: cypher, artifact, quirk)

Getting cypher dangers uses
!gen cydanger <charcyphers> <maxcyphers>
e.g. !gen cydanger 4 5
for a character with 4 cyphers and a max of 5.

TODO: Register commands instead of using string parsing

#Cheatsheet Sources
The numenera cheatsheet was collated by Justin Alexander at http://www.thealexandrian.net
