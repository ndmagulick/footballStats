# footballStats

Pardon my dust and disorginization, this is a work in progress, and shouldn't be treated as a final product.

This is just a small project for me to learn GraphQL, and React. As well as for me to expand my knowledge in Go.

The objective here is to make a website that is similar to "TransferMarkt", but make it look more modern, and quicker.

Plan:
1. Create basic app for just standard domestic leagues 
2. Expand app to handle other tournaments such as domestic cups (DFB Pokal, FA Cup, etc), and continental competitions (Champions League, Europa League, etc)
3. Expand app to handle national teams, and international team matches (friendlies, world cup, euros, etc)

Things I need to do/address:
- Find out why some graphql structs/models are behaving differently from others
- Maybe break out resolver code into their schema's corresponding Go files
- Try to simplify database functions
- Field validation for new/updated records
- Dependency checks for deletions (this will be very important for when we start implementing stats and such)
- Proper error handling
- Figure out how to deal with dates in GraphQL
- Add match stats table
- Seed data
