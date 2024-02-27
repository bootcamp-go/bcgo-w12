# Wave 12

# Frequently Asked Questions (FAQ)

---

## ▶ Github
### Authentication

---

**✔ SSH**
> create new ssh key

```bash
# go to root folder
cd ~
# go to ssh folder
mkdir .ssh && cd .ssh

# save new ssh key
ssh-keygen -t rsa -b 4096 -C "your_email@example.com"
```

> config ssh config file
```bash
touch config
```
open the config file and add the following
```bash
Host *
  AddKeysToAgent yes
  UseKeychain yes
  IdentityFile ~/.ssh/id_rsa
```

---

**✔ Github**
> copy the SSH key
```bash
pbcopy < ~/.ssh/id_rsa.pub
```

> go to your Github account and add the SSH key

---

**✔ Test**

We test the SSH connection to Github (the identity file is specified in the config file)

> test the SSH connection
```bash
ssh -T git@github.com
```

In order to test other identities, you can use the following command
```bash
ssh -i ~/.ssh/key -T git@github.com
```

<br>

Up to now you have configured the SSH key (public / private), indicate ssh which identity file to use and added the SSH key (public) to your Github account.
We can use an optional service to manage the SSH keys (cache) and avoid to enter the phrase each time we use the SSH key.

---

**✔ SSH-Agent - Optional**

initialize ssh-agent as a service of ssh keys management (cache) (optional)

> start the ssh-agent in the background
```bash
eval "$(ssh-agent -s)"
```

> add your SSH private key to the ssh-agent
```bash
ssh-add -K ~/.ssh/id_rsa
```

<br>
<br>
<br>

## ▶ Mercado Libre - SSH Set Up
**✔ Generate SSH Keys**
> go to .ssh folder
```bash
cd ~/.ssh
```

> generate a SSH key's for Mercado Libre
Mercado Libre
```bash
ssh-keygen -t rsa -b 4096 -C "johndoe@mercadolibre.com"  # name of the key is mercadolibre
```

Melisource
```bash
ssh-keygen -t rsa -b 4096 -C "..."  # name of the key is melisource
```

> add identity file to the ssh config file
```bash
Host github.com
  HostName github.com
  User git
  IdentityFile ~/.ssh/mercadolibre
  AddKeysToAgent yes
  UseKeychain yes # macOS
  
Host github.com-emu
  HostName github.com
  User git
  IdentityFile ~/.ssh/melisource
  AddKeysToAgent yes
  UseKeychain yes # macOS
```

---

**✔ Add Keys to Github**
> Go to your mercadolibre Github account and add the SSH key
```bash
pbcopy < ~/.ssh/mercadolibre.pub
```

> Go to your melisource Github account and add the SSH key
```bash
pbcopy < ~/.ssh/melisource.pub
```

---

**✔ Test**

We test the SSH connection to Github (the identity file is specified in the config file)

> test the SSH connection for the mercadolibre account
```bash
ssh -T git@github.com  # it should return "Hi username! ..." with the username of the mercadolibre account
```

> test the SSH connection for the melisource account
```bash 
ssh -T git@github.com-emu  # it should return "Hi username! ..." with the username of the melisource account
```

## ▶ Terminal
in progress...

<br>
<br>
<br>

## ▶ Go
in progress...