About Conv 2.0
==========

## Generale
Questo programma effetua il subnetting di una rete.
Tutte le informazioni necessarie al programma vengono passate tramite un file.
La rete da dividere viene data in input in formato decimale puntato.
Le dimensioni delle sottoreti da ricavare, vengono passate come maschere CIDR.
Le sottoreti vengono allineate e viene ottimizzato lo spazio riempiendo i pool di IP non ancora associati ad alcuna sottorete.

### Comandi
La prima riga del file contiene la rete in formato CIDR che si vuole partizionare:
<p><code>X.X.X.X/Z</code></p>
<p>
    <code>X.X.X.X</code>: è l'indirizzo IPv4 in notazione decimale puntata.
    <br><code>Z</code>: è la maschera CIDR rappresentabile in 3 differenti formati.
</p>
Nelle righe successive vengono indicate le dimensioni delle sottoreti da ricavare. l'ordine di scrittura determina l'ordine in cui verrà ricavato lo spazio di IP.

### Formati delle maschere
<code>Z</code>: Numero di bit destinati alla parte di net ID (formato CIDR standard)
<br><code>-Z</code>: Numero di bit destinati alla parte host ID
<br><code>#Z</code>: Nuemro di indirizzi massimo che deve contenere la rete (deve comprendere Base Address, Brodcast e Getaway)

#### Esempio di file dato in input (prova.txt)
```txt
10.100.30.0/24
#10
27
-3
#5
```

### Lanciare il programma
Il programma viene lanciato  da terminale passando come argomento il nome e il path del file contenente le indicazioni della rete
```cmd
conv.exe prova.txt
```

### Output del programma
L'output avviene sul file <code>a.txt</code>. Il file conterrà gli indirizzi delle sottoreti e tutti i parametri ad esse associati. La prima rete della lista è la rete da dividere.

#### Esempio:
```txt
BaseAddres:	10.100.30.0/24
BroadCast:	10.100.30.255
Getaway:	10.100.30.254
Primo IP:	10.100.30.1
Ultimo IP:	10.100.30.253
Net Mask:	255.255.255.0
Wildcard:	0.0.0.255

#10	--> x.x.x.x/28
BaseAddres:	10.100.30.0/28
BroadCast:	10.100.30.15
Getaway:	10.100.30.14
Primo IP:	10.100.30.1
Ultimo IP:	10.100.30.13
Net Mask:	255.255.255.240
Wildcard:	0.0.0.15

27	--> x.x.x.x/27
BaseAddres:	10.100.30.32/27
BroadCast:	10.100.30.63
Getaway:	10.100.30.62
Primo IP:	10.100.30.33
Ultimo IP:	10.100.30.61
Net Mask:	255.255.255.224
Wildcard:	0.0.0.31

-3	--> x.x.x.x/29
BaseAddres:	10.100.30.16/29
BroadCast:	10.100.30.23
Getaway:	10.100.30.22
Primo IP:	10.100.30.17
Ultimo IP:	10.100.30.21
Net Mask:	255.255.255.248
Wildcard:	0.0.0.7

#5	--> x.x.x.x/29
BaseAddres:	10.100.30.24/29
BroadCast:	10.100.30.31
Getaway:	10.100.30.30
Primo IP:	10.100.30.25
Ultimo IP:	10.100.30.29
Net Mask:	255.255.255.248
Wildcard:	0.0.0.7
```