#!/bin/bash
set -x
echo "install AIDER"
curl -LsSf https://aider.chat/install.sh | sh

create_aider() {
    echo <<'END'
#!/bin/bash
# Aider will take note of all the comments that start or end with AI.
# Comments that include AI! with an exclamation point or AI? with a question mark are special.
# They trigger aider to take action to collect all the AI comments and use them as your instructions.

# AI! triggers aider to make changes to your code.
# AI? triggers aider to answer your question.

# run aider for ide usage
# https://dev.to/iamrj846/connecting-to-the-host-machines-localhost-from-a-docker-container-a-practical-guide-nc4
OLLAMA_API_BASE=http://host.docker.internal:11434 aider --model ollama_chat/deepseek-r1:32b --watch-files
END
}

# Create the aider-ide script with sudon
mkdir -p $HOME/.local/bin
create_aider >$HOME/.local/bin/aider-ide
chmod a+x $HOME/.local/bin/aider-ide

go install golang.org/x/tools/cmd/godoc@latest
