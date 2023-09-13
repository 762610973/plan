let x, y, z;
x = 22;
y = 11;
z = '00';
let a = x+y;
document.getElementById("demo").innerHTML =
    "document.getElementById";

function myFunc1() {
    document.getElementById("demo").innerHTML = "hello function";
}
// 单行注释
/*
* 多行注释
* */

let catName = "carName"
let aa = 1,
    bb = 2,
    cc = 3;

// undefined
let un;

// 常量, 声明时赋值
const PI = 3.141592653589793;

const car = {
    instance:"apple",
    model:"fruit",
    color:"red",
};
// 可以更改属性
car.color = "deep red"
// 可以添加属性
car.owner = "self"

let arr1 = [1,2,3,4];
arr1.push(5);

// 声明常量数组不会使元素不可更改