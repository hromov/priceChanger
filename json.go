package main

var averageHight = `{
	"and": [
	  {
		"<=": [
		  {
			"var": ".height"
		  },
		  55
		]
	  },
	  {
		">=": [
		  {
			"var": ".height"
		  },
		  0
		]
	  }
	]
  }`

var increasedHight = `{
	"and": [
	  {
		"<=": [
		  {
			"var": ".height"
		  },
		  100
		]
	  },
	  {
		">=": [
		  {
			"var": ".height"
		  },
		  55
		]
	  }
	]
  }`

// var increasedHight = `{
// 	"and": [
// 	  {
// 		">": [
// 		  {
// 			"var": ".height"
// 		  },
// 		  100
// 		]
// 	  },
// 	  {
// 		"<": [
// 		  {
// 			"var": ".height"
// 		  },
// 		  200
// 		]
// 	  }
// 	]
//   }`
