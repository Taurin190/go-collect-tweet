var users = [{
    user: 'root',
    pwd: 'root',
    roles: [{
        role: 'dbOwner',
        db: 'twitter'
    }]
}];

for (var i = 0, length = users.length; i < length; ++i) {
    db.createUser(users[i]);
}