# Random Number Generator for Auth

This was specifically asked by someone, so if it's not you then here's what it's for.

**Sorry, I uploaded it to a different account**
> I'll work with what you have so far... <br>
> Looks like ChatGPT to me though...tsk, tsk, tsk
> Becareful with it, because it misses a lot and will have you chasing constantly

The crypto/rand package is excellent!

### Purpose
- Creates a random password hasher.
- Create a public and private hash.
- Initial hash is added/replaced/appended to another hash from a different file. This was specifically requested.

This is purposely simple and generic for a specific person's interest and use case.
Left to do:
- send to database
- configure what you want to add beyond password if you want to create a stronger password
  - use OS
  - use time.