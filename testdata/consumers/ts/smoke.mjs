import {
  SearchResultDisplayBehavior,
  UserService,
} from "@chill-institute/contracts/chill/v4/api_pb";

if (UserService.typeName !== "chill.v4.UserService") {
  throw new Error(`unexpected service type name: ${UserService.typeName}`);
}

if (SearchResultDisplayBehavior.FASTEST !== 2) {
  throw new Error("unexpected enum export value for FASTEST display behavior");
}

console.log("ts consumer import smoke passed");
