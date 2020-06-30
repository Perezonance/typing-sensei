This service is responsible for generating typing tests. This includes a few types of tests:
- Random Typing Test (generated from a lexicon containing the 10,000 most commonly typed words) Options:
    - (Default) no capitalization, punctuation, or numbers
    - Capitalization included
    - Punctuation included
    - Numbers included
    - Advanced chars included
- Lexicon Based Random Typing Test
    - Specifying a specific lexicon which already exists in the DB
- Exact Typing Test
    - Given a hash code we can generate the exact same test
    
Lexicons are entities defined by the set of words that  can be used within a typing test. The 
reason for implementing lexicons is to give typing test catered to more specific fields such
as law, medicine, or even programming. By default, there is a pool of lexicons that users can
choose from which range from some prepared by our team as well as from a list of highly rated
lexicons in the community. 
    