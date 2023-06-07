## How-to

Open the Copilot Chat and ask the following (including the Mermaid diagram):

Generate a daemon program in Python from this mermaid state diagram.
Each state must be a function that returns the next state. 
If multiple next states are possible, choose based on user input.
The daemon should stop after the Finalizing state.

```mermaid
stateDiagram-v2
    state "Loading data" as s2
    state "Processing data" as s3

    [*] --> Initializing      : Daemon started
    Initializing --> s2   : Initialization complete
    s2 --> s3         : Data loaded
    s3 --> s2         : Still some data
    s2 --> Finalizing       : Error
    s3 --> Finalizing        : No more data
    s3 --> Finalizing        : Error
    Finalizing --> [*]           : Daemon stopped
```