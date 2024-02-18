app.get('/resources/:id/detail', function(req, res) {
    let resourceIndex = resources.findIndex(resource => resource.id === parseInt(req.params.id));
    if (resourceIndex !== -1) {
        res.json({
            message: "Resource Details",
            data: resources[resourceIndex]
        });
    } else {
        res.status(404).json({
            message: "No Resource Found with ID " + req.params.id,
            data: {}
        });
    }
});

app.get('/resources/:id', function(req, res) {
    let resourceIndex = resources.findIndex(resource => resource.id === parseInt(req.params.id));
    if (resourceIndex !== -1) {
        res.json({
            message: "Resource Found",
            data: resources[resourceIndex]
        });
    } else {
        res.status(404).json({
            message: "No Resource Found with ID " + req.params.id,
            data: {}
        });
    }
});