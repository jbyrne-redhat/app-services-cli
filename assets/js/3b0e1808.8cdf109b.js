"use strict";(self.webpackChunkwebsite=self.webpackChunkwebsite||[]).push([[9995],{3905:function(e,a,r){r.d(a,{Zo:function(){return l},kt:function(){return f}});var n=r(7294);function t(e,a,r){return a in e?Object.defineProperty(e,a,{value:r,enumerable:!0,configurable:!0,writable:!0}):e[a]=r,e}function s(e,a){var r=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);a&&(n=n.filter((function(a){return Object.getOwnPropertyDescriptor(e,a).enumerable}))),r.push.apply(r,n)}return r}function c(e){for(var a=1;a<arguments.length;a++){var r=null!=arguments[a]?arguments[a]:{};a%2?s(Object(r),!0).forEach((function(a){t(e,a,r[a])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(r)):s(Object(r)).forEach((function(a){Object.defineProperty(e,a,Object.getOwnPropertyDescriptor(r,a))}))}return e}function o(e,a){if(null==e)return{};var r,n,t=function(e,a){if(null==e)return{};var r,n,t={},s=Object.keys(e);for(n=0;n<s.length;n++)r=s[n],a.indexOf(r)>=0||(t[r]=e[r]);return t}(e,a);if(Object.getOwnPropertySymbols){var s=Object.getOwnPropertySymbols(e);for(n=0;n<s.length;n++)r=s[n],a.indexOf(r)>=0||Object.prototype.propertyIsEnumerable.call(e,r)&&(t[r]=e[r])}return t}var i=n.createContext({}),p=function(e){var a=n.useContext(i),r=a;return e&&(r="function"==typeof e?e(a):c(c({},a),e)),r},l=function(e){var a=p(e.components);return n.createElement(i.Provider,{value:a},e.children)},u={inlineCode:"code",wrapper:function(e){var a=e.children;return n.createElement(n.Fragment,{},a)}},m=n.forwardRef((function(e,a){var r=e.components,t=e.mdxType,s=e.originalType,i=e.parentName,l=o(e,["components","mdxType","originalType","parentName"]),m=p(r),f=t,d=m["".concat(i,".").concat(f)]||m[f]||u[f]||s;return r?n.createElement(d,c(c({ref:a},l),{},{components:r})):n.createElement(d,c({ref:a},l))}));function f(e,a){var r=arguments,t=a&&a.mdxType;if("string"==typeof e||t){var s=r.length,c=new Array(s);c[0]=m;var o={};for(var i in a)hasOwnProperty.call(a,i)&&(o[i]=a[i]);o.originalType=e,o.mdxType="string"==typeof e?e:t,c[1]=o;for(var p=2;p<s;p++)c[p]=r[p];return n.createElement.apply(null,c)}return n.createElement.apply(null,r)}m.displayName="MDXCreateElement"},8687:function(e,a,r){r.r(a),r.d(a,{frontMatter:function(){return o},contentTitle:function(){return i},metadata:function(){return p},toc:function(){return l},default:function(){return m}});var n=r(7462),t=r(3366),s=(r(7294),r(3905)),c=["components"],o={},i=void 0,p={unversionedId:"commands/rhoas_kafka_acl_grant-access",id:"commands/rhoas_kafka_acl_grant-access",isDocsHomePage:!1,title:"rhoas_kafka_acl_grant-access",description:"rhoas kafka acl grant-access",source:"@site/docs/commands/rhoas_kafka_acl_grant-access.md",sourceDirName:"commands",slug:"/commands/rhoas_kafka_acl_grant-access",permalink:"/app-services-cli/commands/rhoas_kafka_acl_grant-access",editUrl:"https://github.com/redhat-developer/app-services-cli/docs/commands/rhoas_kafka_acl_grant-access.md",tags:[],version:"current",frontMatter:{},sidebar:"tutorialSidebar",previous:{title:"rhoas_kafka_acl_delete",permalink:"/app-services-cli/commands/rhoas_kafka_acl_delete"},next:{title:"rhoas_kafka_acl_grant-admin",permalink:"/app-services-cli/commands/rhoas_kafka_acl_grant-admin"}},l=[{value:"rhoas kafka acl grant-access",id:"rhoas-kafka-acl-grant-access",children:[{value:"Synopsis",id:"synopsis",children:[]},{value:"Examples",id:"examples",children:[]},{value:"Options",id:"options",children:[]},{value:"Options inherited from parent commands",id:"options-inherited-from-parent-commands",children:[]},{value:"SEE ALSO",id:"see-also",children:[]}]}],u={toc:l};function m(e){var a=e.components,r=(0,t.Z)(e,c);return(0,s.kt)("wrapper",(0,n.Z)({},u,r,{components:a,mdxType:"MDXLayout"}),(0,s.kt)("h2",{id:"rhoas-kafka-acl-grant-access"},"rhoas kafka acl grant-access"),(0,s.kt)("p",null,"Add ACL rules to grant users access to produce and consume from topics"),(0,s.kt)("h3",{id:"synopsis"},"Synopsis"),(0,s.kt)("p",null,"Create Access Control List (ACL) rules that grant the specified user access to produce and consume from topics."),(0,s.kt)("pre",null,(0,s.kt)("code",{parentName:"pre"},"rhoas kafka acl grant-access [flags]\n")),(0,s.kt)("h3",{id:"examples"},"Examples"),(0,s.kt)("pre",null,(0,s.kt)("code",{parentName:"pre"},'# Grant access to principal for consuming messages from all topics\n$ rhoas kafka acl grant-access --consumer --user user_name --topic all --group all\n\n# Grant access to principal for consuming messages from all topics in a specified instance\n$ rhoas kafka acl grant-access --consumer --user user_name --topic all --group all --instance-id c5hv7iru4an1g84pogp0\n\n# Grant access to principal for producing messages to all topics\n$ rhoas kafka acl grant-access --producer --user user_name --topic all\n\n# Grant access to principal for consuming messages from topics starting with "abc"\n$ rhoas kafka acl grant-access --consumer --user user_name --topic-prefix "abc" --group my-group\n\n# Grant access to principal for producing messages to topics starting with "abc"\n$ rhoas kafka acl grant-access --producer --user user_name --topic-prefix "abc"\n\n# Grant access to all users for consuming messages from topic "my-topic"\n$ rhoas kafka acl grant-access --consumer --all-accounts --topic my-topic --group my-group\n\n# Grant access to all users for producing messages to topic "my-topic"\n$ rhoas kafka acl grant-access --producer --all-accounts --topic my-topic\n\n# Grant access to principal for produce and consume messages from all topics\n$ rhoas kafka acl grant-access --producer --consumer --user user_name --topic all --group all\n\n')),(0,s.kt)("h3",{id:"options"},"Options"),(0,s.kt)("pre",null,(0,s.kt)("code",{parentName:"pre"},"      --all-accounts             Set the ACL principal to match all principals (users and service accounts)\n      --consumer                 Add ACL rules that grant the specified principal access to consume messages from topics\n      --group string             Consumer group ID to define ACL rules for\n      --group-prefix string      Prefix name for groups to be selected\n      --instance-id string       Kafka instance ID. Uses the current instance if not set\n      --producer                 Add ACL rules that grant the specified principal access to produce messages to topics\n      --service-account string   Service account client ID used as principal for this operation\n      --topic string             Topic name to define ACL rules for\n      --topic-prefix string      Prefix name for topics to be selected\n      --user string              User ID to be used as principal\n  -y, --yes                      Skip confirmation of this action \n")),(0,s.kt)("h3",{id:"options-inherited-from-parent-commands"},"Options inherited from parent commands"),(0,s.kt)("pre",null,(0,s.kt)("code",{parentName:"pre"},"  -h, --help      Show help for a command\n  -v, --verbose   Enable verbose mode\n")),(0,s.kt)("h3",{id:"see-also"},"SEE ALSO"),(0,s.kt)("ul",null,(0,s.kt)("li",{parentName:"ul"},(0,s.kt)("a",{parentName:"li",href:"/app-services-cli/commands/rhoas_kafka_acl"},"rhoas kafka acl"),"\t - Manage Kafka ACLs for users and service accounts")))}m.isMDXComponent=!0}}]);