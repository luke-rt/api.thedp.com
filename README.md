# The Daily Pennsylvanian API

Docs located [here](https://github.com/luke-rt/api.thedp.com/wiki)

TODO:
- Get dates, sorting by dates working. currently MongoDB stores them as non-standard formats. maybe convert to int?
- MongoDB sometimes stores something as either type A or an empty array, which messes with the type checking. ie Metadata.Value, and DominantMedia. Standardize it in MongoDB
