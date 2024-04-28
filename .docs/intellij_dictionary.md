[Link](https://stackoverflow.com/questions/28147162/how-to-create-a-shared-editable-dictionary-for-intellij-idea-spellchecking#comment44677627_28151692)

Note that, even though the words added to the dictionary by each developer are stored in a separate file, IntelliJ IDEA takes the words added by all developers, and does not highlight any of them as typos. The storage in separate files was designed specifically to avoid merge conflicts when storing the dictionary in a version control system.
Therefore, the default functionality of IntelliJ IDEA should meet your requirements.