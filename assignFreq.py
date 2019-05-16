import networkx as nx

inputFile = open("graph.plain", 'rb')
G = nx.read_edgelist(inputFile)

map = (list(nx.connected_components(G)))

i = 1
for item in map:
    print("BTS-y numer: " + str(item)+ " mają przydzieloną częstotliwość " + str(i))
    i = i + 1

inputFile.close()