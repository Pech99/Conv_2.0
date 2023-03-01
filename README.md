About Conv 2.0
==========

## Generale
Questo programma restituisce informazioni riguardo ad un indirizzo di rete ed opera il subnetting.
Tutte le informazioni necessarie al programma vengono passate tramite un riga di comando dottoforma di argomenti.
La rete di cui ottenere le informazioni e da dividere viene passata come primo argomento in fomrato CIDR (A.B.C.D/M).
Mentre le maschere delle sottoreti da ricavare è possibile specificarle in 3 differenti formati.
Le sottoreti vengono allineate e viene ottimizzato lo spazio riempiendo i pool di IP non ancora associati ad alcuna sottorete.
Al termine della suddivisione vengono restituiti i pool ancora liberi.

### Comandi
<p><code>./conv A.B.C.D/M [M | -M | #M]...</code></p>

<code>A.B.C.D/M</code> : è l'indirizzo IPv4 in notazione decimale puntata formato CIDR.
<br><code>M</code> : Numero di bit destinati alla parte di netID (formato CIDR standard).
<br><code>-M</code> : Numero di bit destinati alla parte hostID.
<br><code>#M</code> : Numero di indirizzi minimo che la rete deve contenere (compreso il Getaway).

#### Esempio di input:
```
./conv 10.100.30.0/24 #10 #30 16 -3 30
```
#### Esempio output:
```
Indirizzo Base della Rete (10.100.30.0/24)
BaseAddres:     10.100.30.0/24
BroadCast:      10.100.30.255
Getaway:        10.100.30.254
Primo IP:       10.100.30.1
Ultimo IP:      10.100.30.253
Net Mask:       255.255.255.0
Wildcard:       0.0.0.255

#10     --> A.B.C.D/28
BaseAddres:     10.100.30.0/28
BroadCast:      10.100.30.15
Getaway:        10.100.30.14
Primo IP:       10.100.30.1
Ultimo IP:      10.100.30.13
Net Mask:       255.255.255.240
Wildcard:       0.0.0.15

#30     --> A.B.C.D/27
BaseAddres:     10.100.30.32/27
BroadCast:      10.100.30.63
Getaway:        10.100.30.62
Primo IP:       10.100.30.33
Ultimo IP:      10.100.30.61
Net Mask:       255.255.255.224
Wildcard:       0.0.0.31

16      --> A.B.C.D/16
Spazio insufficiente per questa rete

-3      --> A.B.C.D/29
BaseAddres:     10.100.30.16/29
BroadCast:      10.100.30.23
Getaway:        10.100.30.22
Primo IP:       10.100.30.17
Ultimo IP:      10.100.30.21
Net Mask:       255.255.255.248
Wildcard:       0.0.0.7

30      --> A.B.C.D/30
BaseAddres:     10.100.30.24/30
BroadCast:      10.100.30.27
Getaway:        -
Primo IP:       10.100.30.25
Secondo IP:     10.100.30.26
Net Mask:       255.255.255.252
Wildcard:       0.0.0.3

Spazzi Rimanenti:
10.100.30.64/26
10.100.30.128/25
10.100.30.28/30

```