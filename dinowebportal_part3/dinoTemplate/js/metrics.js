
$('#dinometrics tr').each(function() {
    var dinohr = $(this).find("td.dinohr");    
    var dinobp = $(this).find("td.dinobp");
    var dinoname = $(this).find("td.dinoname");
    //dinodata.text(dinoname.text());
    
   // if(socket==null)
   var Socket = new WebSocket("ws://localhost:8181/dinodatafeed");
    Socket.onopen = function(event){
        Socket.send(dinoname.text());
        Socket.onmessage = function(event){
            var msg = JSON.parse(event.data);
            dinohr.text(msg.Heartrate);
            dinobp.text(msg.Bloodpressure);
        };
    };
 
 });