"use strict";(self.webpackChunkwebsite=self.webpackChunkwebsite||[]).push([[6947],{3905:function(e,r,t){t.d(r,{Zo:function(){return p},kt:function(){return d}});var a=t(7294);function n(e,r,t){return r in e?Object.defineProperty(e,r,{value:t,enumerable:!0,configurable:!0,writable:!0}):e[r]=t,e}function o(e,r){var t=Object.keys(e);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);r&&(a=a.filter((function(r){return Object.getOwnPropertyDescriptor(e,r).enumerable}))),t.push.apply(t,a)}return t}function c(e){for(var r=1;r<arguments.length;r++){var t=null!=arguments[r]?arguments[r]:{};r%2?o(Object(t),!0).forEach((function(r){n(e,r,t[r])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(t)):o(Object(t)).forEach((function(r){Object.defineProperty(e,r,Object.getOwnPropertyDescriptor(t,r))}))}return e}function s(e,r){if(null==e)return{};var t,a,n=function(e,r){if(null==e)return{};var t,a,n={},o=Object.keys(e);for(a=0;a<o.length;a++)t=o[a],r.indexOf(t)>=0||(n[t]=e[t]);return n}(e,r);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);for(a=0;a<o.length;a++)t=o[a],r.indexOf(t)>=0||Object.prototype.propertyIsEnumerable.call(e,t)&&(n[t]=e[t])}return n}var i=a.createContext({}),l=function(e){var r=a.useContext(i),t=r;return e&&(t="function"==typeof e?e(r):c(c({},r),e)),t},p=function(e){var r=l(e.components);return a.createElement(i.Provider,{value:r},e.children)},u={inlineCode:"code",wrapper:function(e){var r=e.children;return a.createElement(a.Fragment,{},r)}},f=a.forwardRef((function(e,r){var t=e.components,n=e.mdxType,o=e.originalType,i=e.parentName,p=s(e,["components","mdxType","originalType","parentName"]),f=l(t),d=n,m=f["".concat(i,".").concat(d)]||f[d]||u[d]||o;return t?a.createElement(m,c(c({ref:r},p),{},{components:t})):a.createElement(m,c({ref:r},p))}));function d(e,r){var t=arguments,n=r&&r.mdxType;if("string"==typeof e||n){var o=t.length,c=new Array(o);c[0]=f;var s={};for(var i in r)hasOwnProperty.call(r,i)&&(s[i]=r[i]);s.originalType=e,s.mdxType="string"==typeof e?e:n,c[1]=s;for(var l=2;l<o;l++)c[l]=t[l];return a.createElement.apply(null,c)}return a.createElement.apply(null,t)}f.displayName="MDXCreateElement"},1390:function(e,r,t){t.r(r),t.d(r,{frontMatter:function(){return s},contentTitle:function(){return i},metadata:function(){return l},toc:function(){return p},default:function(){return f}});var a=t(7462),n=t(3366),o=(t(7294),t(3905)),c=["components"],s={},i=void 0,l={unversionedId:"commands/rhoas_kafka_acl_create",id:"commands/rhoas_kafka_acl_create",isDocsHomePage:!1,title:"rhoas_kafka_acl_create",description:"rhoas kafka acl create",source:"@site/docs/commands/rhoas_kafka_acl_create.md",sourceDirName:"commands",slug:"/commands/rhoas_kafka_acl_create",permalink:"/app-services-cli/commands/rhoas_kafka_acl_create",editUrl:"https://github.com/redhat-developer/app-services-cli/docs/commands/rhoas_kafka_acl_create.md",tags:[],version:"current",frontMatter:{},sidebar:"tutorialSidebar",previous:{title:"rhoas_kafka_acl",permalink:"/app-services-cli/commands/rhoas_kafka_acl"},next:{title:"rhoas_kafka_acl_delete",permalink:"/app-services-cli/commands/rhoas_kafka_acl_delete"}},p=[{value:"rhoas kafka acl create",id:"rhoas-kafka-acl-create",children:[{value:"Synopsis",id:"synopsis",children:[]},{value:"Examples",id:"examples",children:[]},{value:"Options",id:"options",children:[]},{value:"Options inherited from parent commands",id:"options-inherited-from-parent-commands",children:[]},{value:"SEE ALSO",id:"see-also",children:[]}]}],u={toc:p};function f(e){var r=e.components,t=(0,n.Z)(e,c);return(0,o.kt)("wrapper",(0,a.Z)({},u,t,{components:r,mdxType:"MDXLayout"}),(0,o.kt)("h2",{id:"rhoas-kafka-acl-create"},"rhoas kafka acl create"),(0,o.kt)("p",null,"Create a Kafka ACL"),(0,o.kt)("h3",{id:"synopsis"},"Synopsis"),(0,o.kt)("p",null,"Create Kafka Access Control List (ACL) rules. A Kafka ACL defines how other user accounts and service accounts can interact with a Kafka instance and its resources."),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre"},"rhoas kafka acl create [flags]\n")),(0,o.kt)("h3",{id:"examples"},"Examples"),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre"},'# Create an ACL for user "dev_user" on all topics\n$ rhoas kafka acl create --operation all --permission allow --topic "*" --user dev_user\n\n# Create an ACL for a service account\n$ rhoas kafka acl create --operation all --permission allow --topic "rhoas" --prefix --service-account "srvc-acct-11924479-43fe-42b4-9676-cf0c9aca81"\n\n# Create an ACL for all users for the consumer group resource\n$ rhoas kafka acl create --operation all --permission allow --group "group-1" --all-accounts\n\n')),(0,o.kt)("h3",{id:"options"},"Options"),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre"},'      --all-accounts              Set the ACL principal to match all principals (users and service accounts)\n      --cluster                   Set the resource type to cluster\n      --group string              Set the consumer group resource. When the --prefix option is also passed, this is used as the consumer group prefix\n      --instance-id string        Kafka instance ID. Uses the current instance if not set\n      --operation string          Set the ACL operation. Choose from: "all", "alter", "alter-configs", "create", "delete", "describe", "describe-configs", "read", "write"\n      --permission string         Set the ACL permission. Choose from: "allow", "deny"\n      --prefix                    Determine if the resource should be exact match or prefix\n      --service-account string    Service account client ID used as principal for this operation\n      --topic string              Set the topic resource. When the --prefix option is also passed, this is used as the topic prefix\n      --transactional-id string   Set the transactional ID resource\n      --user string               User ID to be used as principal\n  -y, --yes                       Skip confirmation of this action \n')),(0,o.kt)("h3",{id:"options-inherited-from-parent-commands"},"Options inherited from parent commands"),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre"},"  -h, --help      Show help for a command\n  -v, --verbose   Enable verbose mode\n")),(0,o.kt)("h3",{id:"see-also"},"SEE ALSO"),(0,o.kt)("ul",null,(0,o.kt)("li",{parentName:"ul"},(0,o.kt)("a",{parentName:"li",href:"/app-services-cli/commands/rhoas_kafka_acl"},"rhoas kafka acl"),"\t - Manage Kafka ACLs for users and service accounts")))}f.isMDXComponent=!0}}]);