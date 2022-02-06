# GenPC

Discord bot for automating Numenera (and eventually other system) tables.

Still in early development, for now commands are as follows:

Numenera Helper Commands

`!info ARG`: general overview of requested feature
ARG values: `bot, cypher, oddity, artifact`

`!gen ARG`: generates a random instance of the requested feature
ARG values: `oddity, quirk, cypher, artifact`

Getting cypher dangers uses
`!gen cydanger CHARCYPHERS MAXCYPHERS`
e.g. `!gen cydanger 4 5`
for a character with 4 cyphers and a max of 5.

`!cs` will provide a link to the pdf cheatsheet
`!cs ARG` will provide an in-discord cheatsheet for a requested feature
`!mcs ARG` will provide a mobile-friendly version of the cheatsheet, if applicable
ARG values: `threshold`

TODO:
cheatsheet:
- Modifying tasks
- Special Rolls
- Major/Minor effects
- GM Intrusions
- damage track
- combat actions

gen:
- (eventually) character

# Cheatsheet Sources
The numenera cheatsheet was collated by Justin Alexander at http://www.thealexandrian.net
