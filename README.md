# Tidy Cobra

Avoid Cobra command options in package scope by arranging for the
command to retrieve named flags from its representation in memory.
Commands are created with a closure that amounts to a constructor
accepting injected collaborators, such as databases or network
clients.
