function searchChampion() {
    var input = document.getElementById("champSearch").value.trim();
    if (input) {
        fetch("http://leaguevis.kaspeti.com/api/champions/" + encodeURIComponent(input.toLowerCase()) + "/supports/ally")
            .then(response => {
                if (!response.ok) {
                    throw new Error("API response was not OK: " + response.statusText);
                }
                return response.json()
            })
            .then(data => {
                initializeChart(data);
            })
            .catch(error => {
                console.error("An error occured while fetching data: ", error);
            });
    } else {
        alert("Please enter a champion name");
    }
}


function initializeChart(data) {
    var chart = echarts.init(document.getElementById("main"));

    var options = {
        title: {
            left: "center",
            text: data["champion"],
            subtext: data["totalPlayed"] + " analyzed games",
            textStyle: {
                fontSize: "35",
                color: "white",
            },
            subtextStyle: {
                color: "white",
            },
        },
        tooltip: {
            formatter: function (info) {
                var value = info.value;
                var winrate = value[1];
                winrate = winrate.toFixed(2);

                return [
                    '<div class="tooltip-title">' + echarts.format.encodeHTML(info.name) + '</div>',
                    'Matches Analyzed: ' + value[0] + '<br>',
                    'Winrate: ' + winrate + '<br>',
                ].join('');
            },
        },
        visualMap: {
            orient: "horizontal",
            left: 'center',
            min: 40,
            max: 60,
            text: ["> 60", "< 40"],
            inRange: {
                color: [data["loseColor"], data["neutralColor"], data["winColor"]],
            },
            textStyle: {
                color: "white",
            },
        },
        controller: false,
        series: [{
            type: "treemap",
            itemStyle: {
                borderColor: "black",
                borderWidth: 1,
            },
            breadcrumb: {
                show: false,
            },
            label: {
                normal: {
                    show: true,
                    textStyle: {
                        color: 'black',
                    },
                },
            },
            data: data["data"],
        }],
    };

    chart.setOption(options);

    window.onresize = function() {
        chart.resize();
    };
}

