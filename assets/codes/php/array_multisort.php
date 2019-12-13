<?php

$arr = [
    [
        'id' => 1,
        'name' => 'name1',
        'sort' => 2
    ],
    [
        'id' => 2,
        'name' => 'name2',
        'sort' => 0
    ],
    [
        'id' => 3,
        'name' => 'name3',
        'sort' => 4
    ],
    [
        'id' => 4,
        'name' => 'name4',
        'sort' => 1
    ],
];

$sorts = array_column($arr,'sort');
array_multisort($sorts,SORT_ASC, $arr);

print_r($sorts);
print_r($arr);

/*

Array
(
    [0] => 0
    [1] => 1
    [2] => 2
    [3] => 4
)
Array
(
    [0] => Array
        (
            [id] => 2
            [name] => name2
            [sort] => 0
        )

    [1] => Array
        (
            [id] => 4
            [name] => name4
            [sort] => 1
        )

    [2] => Array
        (
            [id] => 1
            [name] => name1
            [sort] => 2
        )

    [3] => Array
        (
            [id] => 3
            [name] => name3
            [sort] => 4
        )

)
*/