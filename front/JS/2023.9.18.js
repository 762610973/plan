let x1 = 34.00;
console.log(x1);
let x2 = 24;
console.log(x1+x2);

let x = true;
console.log(x === true);

let cars = ['1','2','3'];
for (let i = 0; i < cars.length; i++) {
    console.log(cars[i]);
}

let person = {
    name: "xl",
    age:22,
    func: function () {
        console.log(this.age,this.name);
    }
};
console.log(person.age);

let n = null;
let n2;

console.log(typeof n2);

function func(arg1, arg2) {
    return arg1+arg2;
}

person.func();

let z = Boolean();