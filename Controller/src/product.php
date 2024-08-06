<?php

class Product
{

    public $Id;
    public $Type;
    public $Brand;
    public $Model;
    public $Color;
    public $Size;
    public $Price;

    public function __construct($Id, $Type, $Brand, $Model, $Color, $Size, $Price)
    {

        $this->Id = $Id;
        $this->Type = $Type;
        $this->Brand = $Brand;
        $this->Model = $Model;
        $this->Color = $Color;
        $this->Size = $Size;
        $this->Price = $Price;
    }

    public function introduce()
    {
        return "This is a $this->Id,  $this->Type, $this->Brand, $this->Model, $this->Color, $this->Size, $this->Price";
    }
}
