
class Graph {

	constructor() {
		this.adj = [];
		console.log('Created a graph');
	}

	addEdge(source, target) {
		if(!this.adj[source]) {
			// add elements to vertex list
			this.adj.push([target]);
			this.adj.push([source]);
		} else {
			this.adj[source].push(target);
			this.adj.push([source]);
		}
	}
}

/*

a -> [ b, c ]
b -> [ a ]
c -> [ a ]

0 -> [ 1, 2 ]
1 -> [ 0 ]
2 -> [ 0 ]


{
	[
		[ 1, 2 ],   <--- 0
		[ 0 ],      <--- 1
		[ 0 ]       <--- 2
	]
}

}
*/

var g = new Graph;
g.addEdge(0, 1);
g.addEdge(0, 2);
g.addEdge(0, 3);
console.log(g);
