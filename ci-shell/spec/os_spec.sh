#shellcheck shell=bash

Describe "Unit Test : " unit
    Include src/load.sh
    Context "os.sh"
        It "_debug_option -d functions to set VERBOSE=1"
            When call _debug_option -d
            The variable "$VERBOSE" should equal 1
        End
        It "all_colors function displays text in colour"
            _debug_option -d #set debug option explicitly
            When call all_colors
            The output should include "${RED}RED${NC}"
        End
        It "_display_time function with 300 returns 5 minutes and 0 seconds"
            When call _display_time 300
            The output should equal "5 minutes and 0 seconds"
        End 
        It "_display_time function with 0 returns 0 seconds"
            When call _display_time 0
            The output should equal "0 seconds"
        End
        It "_display_time function with 40000 returns 11 hours 6 minutes and 40 seconds"
            When call _display_time 40000
            The output should equal "11 hours 6 minutes and 40 seconds"
        End 
        It "_display_time function with 400000 returns 14 days 15 hours 6 minutes and 40 seconds"
            When call _display_time 400000
            The output should equal "4 days 15 hours 6 minutes and 40 seconds"
        End 
  End
End

Describe "Integration Test : " integration
    Include src/load.sh
    Context "os.sh"
        It "raise_error should prompt and exit"
            When run raise_error "Error Occured"
            The status should be failure
            The error should include "Exiting..."
        End
        It "_file_exist LICENSE should pass"
            When call  _file_exist Dockerfile
            The status should be success
            The output should include "Found!"
        End
        It "_file_exist LICENSE should fail"
            When call  _file_exist LICENSE
            The status should be failure
            The error should include "Not Found!"
        End
    End
End