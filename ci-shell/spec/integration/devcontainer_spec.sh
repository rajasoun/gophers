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
        It "_get_devcontainer_json should fail with wrong CONFIG_DIR"
            DEV_CONTAINER_JSON_PATH="/dummy"
            When run _get_devcontainer_json
            The status should be failure
            The error should include "devcontainer.json Not Found"
        End
        It "_get_docker_file_path should return valid Dockerfile path"
            Path docker_file="$DEV_CONTAINER_JSON_PATH"
            When call _get_docker_file_path "$(_get_devcontainer_json_corrected)"
            The status should be success
            The output should include ".devcontainer/Dockerfile"
            The path docker_file should be exist
        End
        It "_get_remote_user should return vscode"
            When call _get_remote_user "$(_get_devcontainer_json_corrected)"
            The status should be success
            The output should include "vscode"
        End
        It "_get_build_args should pass"
            When call _get_build_args "$(_get_devcontainer_json_corrected)"
            The status should be success
            The output should include "--build-arg"
        End
        It "_get_shell should pass"
            When call _get_shell "$(_get_devcontainer_json_corrected)"
            The status should be success
            The output should include "/bin/zsh"
        End
        It "_get_port should pass even if configuration not present"
            When call _get_port "$(_get_devcontainer_json_corrected)"
            The status should be success
            The output should equal ""
        End
        It "_get_envs should pass "
            When call _get_envs "$(_get_devcontainer_json_corrected)"
            The status should be success
            The output should equal ""
        End
        It "_get_mount_points should pass "
            When call _get_mount_points "$(_get_devcontainer_json_corrected)"
            The status should be success
            The output should include "--mount"
        End
        It "_get_devcontainer_json_corrected should be a valid json"
            When call _get_devcontainer_json_corrected "$(_get_devcontainer_json)"
            The status should be success
            The output should include "Dockerfile"
        End  
    End 
End
