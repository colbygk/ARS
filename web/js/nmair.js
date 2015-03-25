(function(angular) {
   'use strict';
   angular.module('nma', ['ngAnimate'])
     .controller('nmaController', ['$scope', function($scope) {
           $scope.items = ['flight', 'seat', 'other'];
               $scope.selection = $scope.items[1];
                 }]);
     })(window.angular);


function initSeating() {
    console.log("initSeating");
		var $ = go.GraphObject.make;  // for conciseness in defining templates
		
		seatingDiagram =
			$(go.Diagram, "seatingDiagram",  // must be the ID or reference to div
			{
				autoScale: go.Diagram.Uniform,
				allowMove: false,
				allowGroup: false,
				allowSelect: false,
				allowHorizontalScroll: false,
				allowVerticalScroll: false,
			});
			
		// define the Node template
		seatingDiagram.nodeTemplate =
			$(go.Part, "Spot",
			{ 
				click: onClickShape,
				toolTip:
					$(go.Adornment, "Auto",
					$(go.Shape, { fill: "#EFEFCC" }),
					$(go.TextBlock, { margin: 4 },
						new go.Binding("text", "label"))
					),
			},
			new go.Binding("click", "click"),
			new go.Binding("location", "locPoint"),
			new go.Binding("text", "text"),  // for sorting, apparently more...
			$(go.Shape, "Card",
			{ 
				fill: "lightgray",
				stroke: "black",
				desiredSize: new go.Size(25, 25),
				minSize: new go.Size(25, 25),
			},
			new go.Binding("figure", "figure"),
			new go.Binding("stroke", "stroke"),
			new go.Binding("fill", "fill"),
			new go.Binding("desiredSize", "size")),
			$(go.TextBlock,
				// the default alignment is go.Spot.Center
				new go.Binding("text", "label"))
			);
		
		generateLayout(0); // generate layout
}
	  
function onClickShape(e, obj) {
		console.log("onClickShape");
		var data = seatingDiagram.model.findNodeDataForKey(obj.part.text);
		var newFill = data.fill === data.oldFill ? "#33B5E5" : data.oldFill;
		seatingDiagram.model.setDataProperty(data, "fill", newFill);
		
		//window.alert("Clicked seat: " + obj.part.text + "\n" + data.fill);
}

function greyOut(index) {
		var data = seatingDiagram.model.findNodeDataForKey(index);
		seatingDiagram.model.setDataProperty(data, "fill", "#C0C0C0");
		seatingDiagram.model.setDataProperty(data, "click", null);
		
		//window.alert("Clicked seat: " + obj.part.text + "\n" + data.fill);
}
	 
	// @param layoutMode  0 == mode 1, else mode 2
function generateLayout(layoutMode) {
		console.log("generateLayout");
		var itemCount;
		var blankRows = new go.List("number");
		if (layoutMode === 0) { 
			itemCount = 84;
			blankRows.add(2);
			blankRows.add(6);
		} else {
			itemCount = 48;
			blankRows.add(2);
		}
		
		var colorIndex = 0;
		var colorList = new go.List("string");
		colorList.add("#EB6E44");
		colorList.add("#FFF5C3");
		colorList.add("#8DCDC1");
		
		var nodeDataArray = [];
		var row = -1;
		var column;
		var columnWrap = 12;
		var paddedItemSize = 28;
		var locPoint;
		
		for (var i = 0; i < itemCount; i++) {
			if (i % columnWrap === 0) { ++row; }
			if (blankRows.contains(row)) { ++row; colorIndex = ++colorIndex % 3; }
			//if (i % 20 === 0) { fill = null; }
			
			column = i % columnWrap;
			locPoint = new go.Point((column * paddedItemSize), (row * paddedItemSize));
					
			nodeDataArray.push({
				key:  i + 1,
				text: (i + 1).toString(),
				label: String.fromCharCode(65 + ((row) % 26)),
				fill: colorList.get(colorIndex),
				oldFill: colorList.get(colorIndex),
				locPoint: locPoint
			});
			
			//console.log(i + ":  " + locPoint + "\n");
		}

		++row;
		for (var i = 0; i < columnWrap; i++) {
			column = i % columnWrap;
			locPoint = new go.Point((column * paddedItemSize), (row * paddedItemSize) + 10);
					
			nodeDataArray.push({
				key:  (itemCount + i + 1),
				text:  (itemCount + i + 1).toString(),
				label: (i + 1).toString(),
				fill: null,
				stroke: "black",
				figure: "Circle",
				oldFill: colorList.get(colorIndex),
				locPoint: locPoint
			});
			
			//console.log(i + ":  " + locPoint + "\n");
		}

		// create a Model that does not know about link or group relationships
		seatingDiagram.model = new go.Model(nodeDataArray);
		seatingDiagram.animationManager.isEnabled = false;

		greyOut(25);
		greyOut(26);
		greyOut(37);
		greyOut(38);

		// resize the div
		resizeDiv(++row, paddedItemSize, columnWrap);
}

function resizeDiv(rows, paddedItemSize, columnWrap) {
    console.log("resizeDiv");
		seatingDiagram.startTransaction("resize div");
		
		var div = seatingDiagram.div;
		var newHeight = (rows * paddedItemSize) + (2 * seatingDiagram.padding.top) + 10;
		var newWidth = (paddedItemSize * columnWrap) + (2 * seatingDiagram.padding.left);

		div.style.width = newWidth.toString() + "px";
		div.style.height = newHeight.toString() + "px";
		seatingDiagram.requestUpdate();

		seatingDiagram.commitTransaction("resize div");
}
	  

window.onload = function () {
    $('#depicker').datetimepicker();
    $('#repicker').datetimepicker();

    $("#depicker").on("dp.change",function (e) {
      $('#repicker').data("DateTimePicker").minDate(e.date);
      });
    $("#repicker").on("dp.change",function (e) {
      $('#depicker').data("DateTimePicker").maxDate(e.date);
      });
   // $('#depicker').data("DateTimePicker").minDate(moment());
   // $('#repicker').data("DateTimePicker").minDate(moment());
    //initSeating();
}

