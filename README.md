# Utils

A CLI tool to make some common personal tasks easier. For now, this is primarly a practice project to learn Go using some real-life problems. Ideally this tool will become some sort of personal "pocket-knife" that helps me automate some day to day tasks - especially since Go allows to build any program as a single static binary so I should be able to run it anywhere I need it without any set-up process.

The first problem I want to tackle is that I want to automate some things I have to do manually when editing content for my hugo website.

## Installation

On Linux I can install the latest version of `utils` by running:

```bash
go install github.com/matkv/utils@latest
```

A config file will be created in ~/.config/utils/config.yaml. Or just use the one from the dotfiles repo.
