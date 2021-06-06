#shellcheck shell=bash

Describe "Intgeration Test : " integration
    Include src/load.sh
    Context "devcontainer.sh"
        It "_get_devcontainer_json should return valid formatted json"
            When call _get_devcontainer_json
            The status should be success
            The output should include "devcontainer.json  Found!"
            The output should include "Dockerfile"
        End
        It "_get_devcontainer_json should when run from /tmp"
            cd /tmp || return 1
            When run _get_devcontainer_json
            The status should be failure
            The error should include "devcontainer.json Not Found"
        End
        # It "_get_docker_file_path should return Dockerfile path"
        #     CONFIG_JSON="$(_get_devcontainer_json)"
        #     When call _get_docker_file_path "$CONFIG_JSON"
        #     The status should be success
        #     The output should include "Dockerfile"
        # End
    End 
End

Describe "Unit Test : " unit
    Include src/load.sh
    Context "devcontainer.sh"
        It "_query_json with valid parameters"
            json='{"fruit":{"name":"apple","color":"green","price":1.20}}'
            When call _query_json "$json" "."
            The status should be success
            The output should include "name"
        End   
        It "_query_json with empty parameters"
            When run _query_json 
            The status should be failure
            The error should include "JSON_STRING is Empty!"
        End  
        It "_query_json with one parameters"
            json='{"fruit":{"name":"apple","color":"green","price":1.20}}'
            When run _query_json "$json"
            The status should be failure
            The error should include "QUERY_STRING is Empty!"
        End  
    End
End