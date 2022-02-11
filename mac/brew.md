# brew

## edit source
```
cd $(brew --repo)
git remote -v
# origin	https://github.com/Homebrew/brew (fetch)
# origin	https://github.com/Homebrew/brew (push)
git remote set-url origin https://mirrors.aliyun.com/homebrew/brew.git
git remote -v
# origin	https://mirrors.aliyun.com/homebrew/brew.git (fetch)
# origin	https://mirrors.aliyun.com/homebrew/brew.git (push)
cd $(brew --repo)/Library/Taps/homebrew/homebrew-core
git remote  -v
# origin	https://github.com/Homebrew/homebrew-core (fetch)
# origin	https://github.com/Homebrew/homebrew-core (push)
git remote set-url origin https://mirrors.aliyun.com/homebrew/homebrew-core.git
git remote -v
# origin	https://mirrors.aliyun.com/homebrew/homebrew-core.git (fetch)
# origin	https://mirrors.aliyun.com/homebrew/homebrew-core.git (push)
brew update
# Already up-to-date.
```