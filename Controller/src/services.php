<?php

include "requestDB.php";
header("Access-Control-Allow-Origin: *");
header("Access-Control-Allow-Headers: Content-Type");


if ($_SERVER['REQUEST_METHOD'] === 'POST') {
    $postData = $_POST;

    $key = key($postData);
    $value = $postData[$key];

    switch ($key) {
        case 'name':
            collect_data_DB($array);
            if ($value == "hepsi") {
                echo json_encode($array);
            } else {
                echo json_encode($array[$value]);
            }
            break;
        case 'city':
            echo json_encode(getCities());
            break;
        default:
            echo json_encode(sendRequest($value, $key));
            break;
    }
} else {
    echo json_encode([
        'status' => 'error',
        'message' => 'Invalid request method.'
    ]);
}
