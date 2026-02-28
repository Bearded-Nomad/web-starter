# 🚀 Checklist Nouveau Projet (Go + Templ + Tailwind)

Ce guide permet de déployer un nouveau site sur le VPS en partant du Starter Kit.

---

## 1. Côté VPS (Préparation)
Connecte-toi en SSH et prépare le dossier et le service pour le nouveau binaire.

### Créer le dossier
~~~
$ sudo mkdir -p /var/www/nouveau-site
$ sudo chown max:max /var/www/nouveau-site
~~~

### Créer le service Systemd
~~~
$ sudo nano /etc/systemd/system/nouveau-site.service
~~~

**Contenu à coller** (pense à changer le **PORT** s'il est déjà pris) :
~~~
[Unit]
Description=Service pour Nouveau Site
After=network.target

[Service]
Type=simple
User=maxless
WorkingDirectory=/var/www/nouveau-site
ExecStart=/var/www/nouveau-site/main
Restart=always
Environment=PORT=8081

[Install]
WantedBy=multi-user.target
~~~

### Activer et autoriser le restart
~~~
$ sudo systemctl daemon-reload
$ sudo systemctl enable nouveau-site
$ sudo systemctl start nouveau-site
~~~
# Autoriser le restart sans mot de passe pour GitHub Actions
~~~
$ sudo visudo
~~~
# Ajoute cette ligne TOUT À LA FIN du fichier :
~~~
maxless ALL=(ALL) NOPASSWD: /usr/bin/systemctl restart nouveau-site
~~~

---

## 2. Côté GitHub (Secrets)
Sur le nouveau dépôt GitHub, va dans Settings > Secrets and variables > Actions et ajoute :

- HOST : 51.91.159.7
- USERNAME : maxless
- SSH_KEY : (Ta clé privée id_ed25519 complète, avec BEGIN et END)

---

## 3. Côté Code (Le Workflow)
Dans ton fichier .github/workflows/deploy.yml, modifie les deux chemins :

1. Target : /var/www/nouveau-site
2. Restart : sudo systemctl restart nouveau-site

---

## 4. Premier Déploiement
Depuis ton terminal local (GoLand) :

~~~
$ git init
$ git remote add origin git@github.com:wnomad/nouveau-site.git
$ git add .
$ git commit -m "Initial deploy"
$ git push -u origin main
~~~

✅ C'est terminé !