package cast_test

import "ztaylor.me/cast"

func MakeDataSlice() []interface{} {
	return []interface{}{
		"hello world",
		"1",
		[]int{20, 19},
		map[interface{}]interface{}{
			"stringer": cast.Stringer("foobar stringer"),
		},
	}
}

func MakeDataMap1() map[string]interface{} {
	return map[string]interface{}{
		"k1":   "v1",
		"k2":   "v2",
		"k3":   "v3",
		"i1":   1,
		"i2":   2,
		"i3":   3,
		"b1":   true,
		"b2":   false,
		"map":  map[string]interface{}{},
		"json": cast.JSON{},
		"dat1": "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
		"dat2": "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		"dat3": "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz",
	}
}

func MakeDataMap2() map[string]interface{} {
	const data = `{"cards":[{"image":"/img/card/1.jpg","name":"Time Runner","text":"","type":"being","powers":[{"usesturn":true,"useskill":false,"id":1,"text":"Draw a Card","costs":{},"trigger":"","target":"self"}],"costs":{"1":1},"body":{"attack":0,"health":1},"id":1},{"image":"/img/card/2.jpg","name":"Ifrit","text":"","type":"being","powers":[{"trigger":"","target":"player-being","usesturn":false,"useskill":false,"id":1,"text":"Deal 1 damage to Target Being or Player","costs":{"2":1}}],"costs":{"2":1},"body":{"attack":1,"health":1},"id":2},{"powers":[],"costs":{"3":1},"body":{"attack":2,"health":2},"id":3,"image":"/img/card/4.jpg","name":"Zealot","text":"So where is the party?","type":"being"},{"name":"Vine Spirit","text":"","type":"being","powers":[{"usesturn":false,"useskill":false,"id":1,"text":"Gain 1 Attack","costs":{},"trigger":"sunrise","target":"self"}],"costs":{"4":1},"body":{"attack":1,"health":1},"id":4,"image":"/img/card/3.jpg"},{"powers":[{"text":"Target Being becomes Asleep","costs":{},"trigger":"","target":"being","usesturn":true,"useskill":false,"id":1}],"costs":{"5":1},"body":{"attack":1,"health":1},"id":5,"image":"/img/card/5.jpg","name":"Water Dancer","text":"","type":"being"},{"type":"being","powers":[{"id":1,"text":"Target Being or Player gains Health equal to my Health","costs":{},"trigger":"","target":"self","usesturn":false,"useskill":true}],"costs":{"6":1},"body":{"attack":1,"health":1},"id":6,"image":"/img/card/6.jpg","name":"Pixie","text":""},{"body":{"attack":0,"health":1},"id":7,"image":"/img/card/7.jpg","name":"Nightmare Ader","text":"","type":"being","powers":[{"costs":{},"trigger":"","target":"being","usesturn":true,"useskill":true,"id":1,"text":"Target Being deals damage to itself"}],"costs":{"7":1}},{"body":{},"id":8,"image":"/img/card/8.jpg","name":"New Element","text":"Prepare for tomorrow","type":"spell","powers":[{"text":"Create a new Element","costs":{},"trigger":"play","target":"self","usesturn":false,"useskill":false,"id":1}],"costs":{"1":1}},{"body":{},"id":9,"image":"/img/card/9.jpg","name":"Burn","text":"Boom! Roasted!","type":"spell","powers":[{"id":1,"text":"Deal 2 damage to Target Being","costs":{},"trigger":"play","target":"being","usesturn":false,"useskill":false}],"costs":{"2":1}},{"powers":[{"id":1,"text":"Target Being or Item becomes Awake","costs":{},"trigger":"play","target":"being-item","usesturn":false,"useskill":false}],"costs":{"3":1},"body":{},"id":10,"image":"img/card/10.jpg","name":"Energize","text":"Now is the time!","type":"spell"},{"text":"The best defense is a strong offense","type":"spell","powers":[{"text":"Target Being gains 1 Attack","costs":{},"trigger":"play","target":"being","usesturn":false,"useskill":false,"id":1}],"costs":{"4":1},"body":{},"id":11,"image":"/img/card/11.jpg","name":"Inspire Growth"},{"name":"Wormhole","text":"It's a trap!","type":"spell","powers":[{"useskill":false,"id":1,"text":"Add Target Being to its owners Hand, and remove it from the Present","costs":{},"trigger":"play","target":"being","usesturn":false}],"costs":{"5":1},"body":{},"id":12,"image":"/img/card/12.jpg"},{"costs":{"6":1},"body":{},"id":13,"image":"/img/card/13.jpg","name":"Grace","text":"Is that better?","type":"spell","powers":[{"trigger":"play","target":"being","usesturn":false,"useskill":false,"id":1,"text":"Target Being gains 2 Health","costs":{}}]},{"body":{},"id":14,"image":"/img/card/14.jpg","name":"Memorialize","text":"Never forget","type":"spell","powers":[{"costs":{},"trigger":"play","target":"mypast-being","usesturn":false,"useskill":false,"id":1,"text":"Add a Card to your Hand that is a copy of a Being Card from your Past"}],"costs":{"7":1}}],"packs":[{"cost":3,"image":"/img/pack/1.jpg","cards":[1,2,3,4,5,6,7],"id":1,"name":"Alpha Beings Pack (1)","size":1},{"cost":5,"image":"/img/pack/2.jpg","cards":[1,2,3,4,5,6,7],"id":2,"name":"Alpha Beings Pack (3)","size":3},{"cards":[1,2,3,4,5,6,7],"id":3,"name":"Alpha Beings Pack (5)","size":5,"cost":7,"image":"/img/pack/3.jpg"},{"cost":3,"image":"/img/pack/4.jpg","cards":[8,9,10,11,12,13,14],"id":4,"name":"Alpha Spells Pack (1)","size":1},{"id":5,"name":"Alpha Spells Pack (3)","size":3,"cost":5,"image":"/img/pack/5.jpg","cards":[8,9,10,11,12,13,14]},{"image":"/img/pack/6.jpg","cards":[8,9,10,11,12,13,14],"id":6,"name":"Alpha Spells Pack (5)","size":5,"cost":7}],"decks":[{"id":1,"name":"AllCards","cover":"/img/card/4.jpg","cards":{"8":3,"10":3,"3":3,"11":3,"1":3,"5":3,"2":3,"9":3,"14":3,"13":3,"12":3,"7":3,"4":3,"6":3}},{"id":2,"name":"WREG","cover":"/img/card/1.jpg","cards":{"3":3,"4":3,"8":3,"9":3,"10":3,"1":3,"2":3}},{"name":"GBVK","cover":"/img/card/7.jpg","cards":{"14":3,"4":3,"5":3,"6":3,"7":3,"12":3,"13":3},"id":3}],"users":25}`
	json := cast.JSON{}
	_ = cast.DecodeJSON(cast.NewReader(data), &json)
	return json
}
