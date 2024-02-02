
class Graph {

	constructor() {
		this.adj = {};
		console.log('Created a graph');
	}

	addEdge(source, target) {
		if(!this.adj[source]) {
			// add elements to vertex list
			this.adj[target] = [source];
			this.adj[source] = [target];
		} else {
			this.adj[source].push(target);
			this.adj[target] = [source];
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

A collection of objects won't work because I need to access a vertex by a text key.

Javascript can create a collection (array) of objects.

[
	{ foo:1, bar:2 },
	{ foo:2, bar:3 },
	{ foo:0, bar:1 }
]

This allows objects to be pushed onto an array (added to the end of the collection).

I'm really pissed at SO people. An associative array is "an unordered collection of name-value pairs." An object in JavaScript is an associative array, an object in Perl is an associative array.

Interestingly, I read that JavaScript arrays are simply objects. An array object has a length property, and methods that operate only on properties with numeric names. So, it's like PHP arrays. Curiously,
arrays are _unordered_, and ordered access is achieved by sequentially addressing the numerically named pairs using an index or counter (for loop). That is, for i=0 to n, does the counting and ordering, and the array is stored as an object.

Why didn't the JS people just say this?

This means an object's property can be accesed by
obj['foo'] = 'bar'

or

obj.foo = bar

The array notation works with alphanumeric keys.

Because in JavaScript object properties are accessed by property name, Object.propertyName, its not possible
to address a property in the same way as associative arrays in other languages, obj.foo is a literal, and I'm
not sure if JavaScript supports using a variable for a property name (I don't see how it could with the dot syntax).

You can say

a['foo'] = 1

Syntactically, this suggests an equivalence of object and array, which is actually the case.

I'm guessing that

for i = 1 to 8 ...
a[key] = 0;

will produce an object/array with numeric keys

so

how to produce an object/array with alphanumeric keys?

The reason that foo[bar].push(qux) fails, is that a generic object does not have the push method defined.

This, this.adj[source].push(target) does work, because it pushes onto an array that is a named property of an object.

var adj = {
	'foo': [ 'bar', 'qyx']
};
adj[foo].push('xyz');


{
	'c1b': [ 'c2b', 'c3b'],
	'c2b': [ 'c1b' ],
	'c3b': [ 'c2b' ]
}
*/

var g = new Graph;
g.addEdge('c0b', 'c1b');
g.addEdge('c0b', 'c2b');
g.addEdge('c0b', 'c3b');
console.log(g);
