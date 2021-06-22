#shellcheck shell=bash

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