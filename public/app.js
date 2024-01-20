function emptyUserForm() {
  return {
    title: "Add User",
    username: "",
    password: "",
    isAdmin: false,
    action: "/admin/accounts/new",
    state: "new",
    new: true,
  };
}

function userForm(id, username, isAdmin) {
  return {
    username: username,
  };
}

function logs() {
  let data = Alpine.reactive({ logstream: [] });
  if (typeof EventSource !== "undefined") {
    let source = new EventSource("/logstream");
    source.onmessage = function (event) {
      let parts = event.data.split(" ");

      data.logstream.push({
        timestamp: parts[0],
        level: parts[1],
        caller: parts[2],
        message: parts.slice(3).join(" "),
      });
    };
  } else {
    data.logstream.push(
      "Sorry, your browser does not support server-sent events..."
    );
  }

  return data;
}
