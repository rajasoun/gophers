#shellcheck shell=bash

Describe "Intgeration Test : "
    Include src/load.sh
    Context "devcontainer.sh"
        # It "__get_git_workspace returns workspace path"
        #     When call _get_workspace
        #     The output should include "workspace"
        # End
    End 
End