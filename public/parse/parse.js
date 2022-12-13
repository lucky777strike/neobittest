function parse() {
    query = document.getElementById('query1').value
    let xhr = new XMLHttpRequest();
    xhr.open("GET", "/api/parse?query=" + query);
    xhr.onreadystatechange = function () {
       if (xhr.readyState === 4) {
          console.log(xhr.responseText);
          if (xhr.responseText =="true"){
            alert("Парсинг запущен");
        }else {
            alert("Ошибка");
          };
       }};
    
    xhr.send();
}



// Builds the HTML Table out of myList.
function buildTable(selector) {


  let xhr = new XMLHttpRequest();
  xhr.open("GET", "/api/getAllTasks");
  xhr.onreadystatechange = function () {
     if (xhr.readyState === 4 && xhr.status === 200) {
        console.log(xhr.responseText);
        myList = JSON.parse(xhr.responseText);
        var columns = addAllColumnHeaders(myList, selector);

        for (var i = 0; i < myList.length; i++) {
          var row$ = $('<tr/>');
          for (var colIndex = 0; colIndex < columns.length; colIndex++) {
            var cellValue = myList[i][columns[colIndex]];
            if (cellValue == null) cellValue = "";
            row$.append($('<td/>').html(cellValue));
          }
          $(selector).append(row$);
        }

     }};
  
  xhr.send();

  
 
}

// Adds a header row to the table and returns the set of columns.
// Need to do union of keys from all records as some records may not contain
// all records.
function addAllColumnHeaders(myList, selector) {
  var columnSet = [];
  var headerTr$ = $('<tr/>');

  for (var i = 0; i < myList.length; i++) {
    var rowHash = myList[i];
    for (var key in rowHash) {
      if ($.inArray(key, columnSet) == -1) {
        columnSet.push(key);
        headerTr$.append($('<th/>').html(key));
      }
    }
  }
  $(selector).append(headerTr$);

  return columnSet;
}

function removeTasks() {
  let xhr = new XMLHttpRequest();
  xhr.open("GET", "/api/removeAllTasks");
  xhr.onreadystatechange = function () {
     if (xhr.readyState === 4) {
        if (xhr.status === 200){
          alert("Задачи удалены");
      }else {
          alert("Ошибка");
        };
     }};
  
  xhr.send();
}
