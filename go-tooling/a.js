var fieldsInsideTypes =[];
var businessObject = {};
businessObject.types = {
    "address": [
        {
            "city": "Pune"
        },
        {
            "city": "Mumbai"
        },
        {
            "city": "Delhi"
        },
        {
            "city": "Chennai"
        }
    ],
    "email": "a@a.in"
};
/*for (var k = 0; k<3; k++){
    businessObject.types.push({
        name: "abc" + k
    });
}*/
for(var p in businessObject.types){
  fieldsInsideTypes.push(businessObject.types[p]);
} 

console.log(fieldsInsideTypes);