function faktorial(n: number): number {
    if (n === 0) return 1;
    let hasil = 1;
    for (let i = 1; i <= n; i++) {
        hasil *= i;
    }
    return hasil;
}

function f(n: number): number {
    const fakt = faktorial(n);
    const pembagi = Math.pow(2, n);
    return Math.ceil(fakt / pembagi);
}

console.log(f(0));


