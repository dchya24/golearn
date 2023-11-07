<?php


// example code
$data = [
    [
        "g" => "G01",
        "m" => ["K01","K10","K12","K13"],
        "belief" => 0.65000,
        "plau" => 0.35000
    ],
    [
        "g" => "G02",
        "m" => ["K01","K12"],
        "belief" => 0.50000,
        "plau" => 0.50000
    ],
    [
        "g" => "G03",
        "m" => ["K01","K12"],
        "belief" => 0.60000,
        "plau" => 0.4
    ]
];

$m = [];

$m[0] = [
    "m" => join(",", $data[0]["m"]),
    "b" => $data[0]["belief"],
    "p" => $data[0]["plau"]
];

$m[1] = [
    "m" => join(",", $data[1]["m"]),
    "b" => $data[1]["belief"],
    "p" => $data[1]["plau"]
];

$a = [
    "m" => join(",", array_intersect(explode(",", $m[0]["m"]), explode(",", $m[1]["m"]))),
    "v" => $m[0]["b"] * $m[1]["b"]
];

$b = [
    "m" => $m[0]["m"],
    "v" => $m[0]["b"] * $m[1]["p"]
];

$c = [
    "m" => $m[1]["m"],
    "v" => $m[1]["b"] * $m[0]["p"]
];


$d = [
    "m" => "p",
    "v" => $m[1]["p"] * $m[0]["p"]
];

$matrix = [
    [$a, $b],
    [$c, $d]
];

$g = [];

for($i = 0; $i < count($matrix); $i++){
    for($j = 0; $j < count($matrix[$i]); $j++){
        $temp = $matrix[$i][$j]['m'];
        if(!array_key_exists($temp, $g)){
            $g["$temp"]["v"] = $matrix[$i][$j]['v'];
            $g["$temp"]["c"] = 1;
        }else{
            $g["$temp"]["v"] += $matrix[$i][$j]['v'];
            $g["$temp"]["c"] += 1;
        }
    }

}

foreach($g as $key => $val){
    $avg = $g[$key]["v"] / $g[$key]["c"];
    $g[$key] = $avg;
}

$m[0] = [
    join(",", $data[0]["m"]) => $data[0]["belief"],
    "p" => $data[0]["belief"]
];

$m[1] = [
    join(",", $data[1]["m"]) => $data[1]["belief"],
    "p" => $data[1]["belief"]
];

$m[2] = $g;

if(count($data) > 2){
    $j = 3;
    for($i = 2 ; $i < count($data); $i++){
        $m[$j] = [
            join(",", $data[$i]["m"]) => $data[$i]["belief"],
            "p" => $data[$i]["plau"]
        ];
        $now = $m[$j];
        $t = $m[$j-1];
        $tmp = [];

        foreach($t as $key1 => $val1){
            foreach($now as $key2 => $val2){
                print_r();
                // if($key2 != "plau" || $key1 != "plau"){
                //     $k = array_intersect(explode(",",$key1), explode(",",$key2));
                //     $tmp = [
                //         join(",", $k), 
                //         $val1 * $val2
                //     ];
                // }
            }
        }

        $m[$tmp];
    }
}

// print_r($m);