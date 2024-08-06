<?php

include "product.php";

function collect_data_DB(&$array)
{
    $curl = curl_init();

    if (!extension_loaded('curl')) {
        echo "Curl Hatasi";
        exit();
    }

    curl_setopt($curl, CURLOPT_URL, "destination service address");
    curl_setopt($curl, CURLOPT_RETURNTRANSFER, true);

    $output = curl_exec($curl);

    if (curl_errno($curl)) {
        echo 'Curl error: ' . curl_errno($curl);
    }

    curl_close($curl);
    $array2 = json_decode($output);

    $array = [
        "shoes" => [],
        "hat" => [],
        "pants" => [],
        "tshirt" => [],
        "shorts" => []
    ];

    foreach ($array2 as $product1) {
        switch ($product1->Type) {
            case "shoes":
                $array["shoes"][] = $product1;
                break;

            case "hat":
                $array["hat"][] = $product1;
                break;

            case "pants":
                $array["pants"][] = $product1;
                break;

            case "tshirt":
                $array["tshirt"][] = $product1;
                break;

            case "shorts":
                $array["shorts"][] = $product1;
                break;

            default:
                // Type bilinmiyorsa bir şey yapma
                break;
        }
    }
}

function sendRequest($data, $name)
{
    $curl = curl_init();
    if (!extension_loaded('curl')) {
        echo "Curl Hatasi";
        exit();
    }

    curl_setopt($curl, CURLOPT_URL, "destination service address" . $name);
    curl_setopt($curl, CURLOPT_POST, 1);
    curl_setopt($curl, CURLOPT_POSTFIELDS, json_encode($data));
    curl_setopt($curl, CURLOPT_RETURNTRANSFER, true);

    $output = curl_exec($curl);

    if (curl_errno($curl)) {
        echo 'Curl error: ' . curl_errno($curl);
    }

    curl_close($curl);
    return json_decode($output);
}
function getCities()
{
    $curl = curl_init();

    if (!extension_loaded('curl')) {
        echo "Curl Hatasi";
        exit();
    }

    curl_setopt($curl, CURLOPT_URL, "destination service address");
    curl_setopt($curl, CURLOPT_RETURNTRANSFER, true);

    $output = curl_exec($curl);
    if (curl_errno($curl)) {
        echo 'Curl error: ' . curl_errno($curl);
    }
    curl_close($curl);
    return json_decode($output);
}
