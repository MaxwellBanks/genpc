# GenPC

Discord bot for automating Numenera (and eventually other system) tables.

Still in early development, for now commands are as follows:

Numenera Helper Commands

`!info ARG`: general overview of requested feature
ARG values: `bot, cypher, oddity, artifact, effects`

`!gen ARG`: generates a random instance of the requested feature
ARG values: `oddity, quirk, cypher, artifact`

Getting cypher dangers uses
`!gen cydanger CHARCYPHERS MAXCYPHERS`
e.g. `!gen cydanger 4 5`
for a character with 4 cyphers and a max of 5.

`!cs`: provides link to cheatsheet
`!cs ARG`: provides in-app cheatsheets for particular features
ARG values: `threshold, effects, rolls`

`!mcs`: functions the same way as !cs, but with truncated tables for mobile users.

`!kronk` will return Kronk's random reaction

TODO:
cheatsheet:
- Modifying tasks
- GM Intrusions
- damage track
- combat actions

gen:
- (eventually) character

# Cheatsheet Sources
The numenera cheatsheet was collated by Justin Alexander at http://www.thealexandrian.net
