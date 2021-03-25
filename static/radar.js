function now()
{
    return typeof window.performance !== 'undefined'
        ? window.performance.now()
        : 0;
}
var CSV = null;
var config = {
    delimiter: "",
    header: true,
    dynamicTyping: false,
    skipEmptyLines: false,
    preview: 0,
    complete: results => {
        end = now();

        if (results && results.errors)
        {
            if (results.errors)
            {
                errorCount = results.errors.length;
                firstError = results.errors[0];
            }
            if (results.data && results.data.length > 0)
                rowCount = results.data.length;
        }
        CSV = results.data[0];
    }
}

function parseCSV(csv) {
    parsed = Papa.parse(csv, config)

    var range = ["a", "b", "c", "d", "e", "f", "g", "h", "i", "j"];
    var grades = 
        [
            ["index", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j"]
        ];

    for (var i = 0; i < range.length; i++) {
        var bigR = range[i].toUpperCase();
        var grade = [range[i]];

        for (var a = 0; a < range.length-1; a++) {  
            if(i==a){
                grade.push(CSV[range[a].toUpperCase()]);
                grade.push(CSV[range[a+1].toUpperCase()]);
            }else{
                grade.push(0);
            }
        }

        grades.push(grade);
        grades[10] = ["j", CSV[range[0].toUpperCase()],0,0,0,0,0,0,0,0,CSV[range[9].toUpperCase()]];
    }
    return grades;
}

var mycfg = {
    color: function(i) {
        c = [
            '#1A1919',
            '#111111',
            '#3E3E3E',
            '#323232',
            '#424141',
            '#605E5E',
            '#7E7D7D',
            '#9E9E9E',
            '#CCCCCC',
            '#0C0C0C',
        ];
        return c[i];
    },
}

function showRadar(parsedcsv) {
    var data = [];
    var chart = RadarChart.chart();
    var w = h = 280;

    headers = []
    parsedcsv.map(function(item, i) {
        if (i == 0) {
            headers = item;
        } else {
            newSeries = {};
            item.map(function(v, j) {
                if (j == 0) {
                    newSeries.className = v;
                    newSeries.axes = [];
                } else {
                    newSeries.axes.push({ "axis": [headers[j]], "value": parseFloat(v) });
                }
            });
            data.push(newSeries);
        }
    })

    RadarChart.defaultConfig.radius = 3;
    RadarChart.defaultConfig.w = w;
    RadarChart.defaultConfig.h = h;

    RadarChart.levelTrick = true;
    RadarChart.draw("#chart-container", data, mycfg);

    return data;
}
var chart_div = document.getElementById('chart-container');

var csv = parseCSV(csv);
var data = showRadar(csv);
