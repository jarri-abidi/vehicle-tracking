PUBLIC_KEY = age1xrn0v8lyg5ddc0r492ccs66jpqx292xq9zlxt3an0r5448nd45qqglndl8 

encrypt:
	sops --encrypt --age $(PUBLIC_KEY) secret.env > .env

decrypt:
	sops --decrypt .env
