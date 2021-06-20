# gophers

Career guide to become a Awesome Go Developer

Project is aimed at creating a study guide for Early in Career developers to learn Go and gain expertise in handling real world projects

# TDD - IaaC - In CI Shell

Open Bash Terminal 

1. `./ci.sh build`

1. `./ci.sh shell`

1. Run `goss --gossfile ci-shell/spec/goss-tests/gossfile.goss.yaml validate` 
to check the container 


# TDD - IaaC - Dev Container 

Open Terminal 

1. Run `shellspec -c ci-shell/spec --tag unit,integration,iaac  --kcov --profile` to check if the container provision is valid

1. Run `goss validate `
