var arr = [{
    department_id: 'D1',
    name: 'D1',
    children: [
        {
            department_id: 'D1-1',
            name: 'D1-1',
        },
        {
            department_id: 'D1-2',
            name: 'D1-2',
            children: [
                {
                    department_id: 'D1-2-1',
                    name: 'D1-2-1',
                },
                {
                    department_id: 'D1-2-2',
                    name: 'D1-2-2',
                },
                {
                    department_id: 'D1-2-3',
                    name: 'D1-2-3',
                    children: [
                        {
                            department_id: 'D1-2-3-1',
                            name: 'D1-2-3-1',
                        },
                        {
                            department_id: 'D1-2-3-2',
                            name: 'D1-2-3-2',
                        },
                    ]
                },
            ]
        }
    ]
}];

var arr2 = ['D1-2-3-2','D1-2-2'];
function test(arr)
{
    for (var i in arr) {
        if ('object' === typeof arr[i]) {
            if (0 <= arr2.indexOf(arr[i].department_id)) {
                arr[i].disabled = true;
            }
            if (undefined !== arr[i].children) {
                test(arr[i].children)
            }
        }
    }
    return arr;
}

arr = test(arr);
console.log(arr);
