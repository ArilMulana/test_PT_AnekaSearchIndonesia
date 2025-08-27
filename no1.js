function faktorial(n) {
    if (n === 0)
        return 1;
    var hasil = 1;
    for (var i = 1; i <= n; i++) {
        hasil *= i;
    }
    return hasil;
}
function f(n) {
    var fakt = faktorial(n);
    var pembagi = Math.pow(2, n);
    return Math.ceil(fakt / pembagi);
}
console.log(f(0));
