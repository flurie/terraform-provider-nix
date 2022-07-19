#InputAddressed: {
	path: string
}

#CAFixed: {
	path: string
	hashAlgo: string
	hash: string
}

#CAFloating: {
	hashAlgo: string
}

#Deferred: {}

#Impure: {
	hashAlgo: string
	impure: true
}

#ListStr: [string]

#Derivation: {
	outputs: { for _, v  }
}

#ShowDerivation: {

}
