"use strict";(self.webpackChunkwebsite=self.webpackChunkwebsite||[]).push([[5562],{3905:function(e,r,n){n.d(r,{Zo:function(){return p},kt:function(){return f}});var a=n(7294);function t(e,r,n){return r in e?Object.defineProperty(e,r,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[r]=n,e}function o(e,r){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);r&&(a=a.filter((function(r){return Object.getOwnPropertyDescriptor(e,r).enumerable}))),n.push.apply(n,a)}return n}function s(e){for(var r=1;r<arguments.length;r++){var n=null!=arguments[r]?arguments[r]:{};r%2?o(Object(n),!0).forEach((function(r){t(e,r,n[r])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):o(Object(n)).forEach((function(r){Object.defineProperty(e,r,Object.getOwnPropertyDescriptor(n,r))}))}return e}function c(e,r){if(null==e)return{};var n,a,t=function(e,r){if(null==e)return{};var n,a,t={},o=Object.keys(e);for(a=0;a<o.length;a++)n=o[a],r.indexOf(n)>=0||(t[n]=e[n]);return t}(e,r);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);for(a=0;a<o.length;a++)n=o[a],r.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(t[n]=e[n])}return t}var i=a.createContext({}),u=function(e){var r=a.useContext(i),n=r;return e&&(n="function"==typeof e?e(r):s(s({},r),e)),n},p=function(e){var r=u(e.components);return a.createElement(i.Provider,{value:r},e.children)},l={inlineCode:"code",wrapper:function(e){var r=e.children;return a.createElement(a.Fragment,{},r)}},m=a.forwardRef((function(e,r){var n=e.components,t=e.mdxType,o=e.originalType,i=e.parentName,p=c(e,["components","mdxType","originalType","parentName"]),m=u(n),f=t,k=m["".concat(i,".").concat(f)]||m[f]||l[f]||o;return n?a.createElement(k,s(s({ref:r},p),{},{components:n})):a.createElement(k,s({ref:r},p))}));function f(e,r){var n=arguments,t=r&&r.mdxType;if("string"==typeof e||t){var o=n.length,s=new Array(o);s[0]=m;var c={};for(var i in r)hasOwnProperty.call(r,i)&&(c[i]=r[i]);c.originalType=e,c.mdxType="string"==typeof e?e:t,s[1]=c;for(var u=2;u<o;u++)s[u]=n[u];return a.createElement.apply(null,s)}return a.createElement.apply(null,n)}m.displayName="MDXCreateElement"},3638:function(e,r,n){n.r(r),n.d(r,{frontMatter:function(){return c},contentTitle:function(){return i},metadata:function(){return u},toc:function(){return p},default:function(){return m}});var a=n(7462),t=n(3366),o=(n(7294),n(3905)),s=["components"],c={},i=void 0,u={unversionedId:"commands/rhoas_kafka_consumer-group",id:"commands/rhoas_kafka_consumer-group",isDocsHomePage:!1,title:"rhoas_kafka_consumer-group",description:"rhoas kafka consumer-group",source:"@site/docs/commands/rhoas_kafka_consumer-group.md",sourceDirName:"commands",slug:"/commands/rhoas_kafka_consumer-group",permalink:"/app-services-cli/commands/rhoas_kafka_consumer-group",editUrl:"https://github.com/redhat-developer/app-services-cli/docs/commands/rhoas_kafka_consumer-group.md",tags:[],version:"current",frontMatter:{},sidebar:"tutorialSidebar",previous:{title:"rhoas_kafka_acl_list",permalink:"/app-services-cli/commands/rhoas_kafka_acl_list"},next:{title:"rhoas_kafka_consumer-group_delete",permalink:"/app-services-cli/commands/rhoas_kafka_consumer-group_delete"}},p=[{value:"rhoas kafka consumer-group",id:"rhoas-kafka-consumer-group",children:[{value:"Synopsis",id:"synopsis",children:[]},{value:"Examples",id:"examples",children:[]},{value:"Options inherited from parent commands",id:"options-inherited-from-parent-commands",children:[]},{value:"SEE ALSO",id:"see-also",children:[]}]}],l={toc:p};function m(e){var r=e.components,n=(0,t.Z)(e,s);return(0,o.kt)("wrapper",(0,a.Z)({},l,n,{components:r,mdxType:"MDXLayout"}),(0,o.kt)("h2",{id:"rhoas-kafka-consumer-group"},"rhoas kafka consumer-group"),(0,o.kt)("p",null,"Describe, list, and delete consumer groups for the current Kafka instance"),(0,o.kt)("h3",{id:"synopsis"},"Synopsis"),(0,o.kt)("p",null,"View and delete consumer groups for the current Kafka instance."),(0,o.kt)("p",null,"These commands operate on the current Kafka instance. To select the Kafka instance, use the \u201crhoas kafka use\u201d command."),(0,o.kt)("h3",{id:"examples"},"Examples"),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre"},"# Delete a consumer group\nrhoas kafka consumer-group delete --id consumer_group_1\n\n# List all consumer groups\nrhoas kafka consumer-group list\n\n")),(0,o.kt)("h3",{id:"options-inherited-from-parent-commands"},"Options inherited from parent commands"),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre"},"  -h, --help      Show help for a command\n  -v, --verbose   Enable verbose mode\n")),(0,o.kt)("h3",{id:"see-also"},"SEE ALSO"),(0,o.kt)("ul",null,(0,o.kt)("li",{parentName:"ul"},(0,o.kt)("a",{parentName:"li",href:"/app-services-cli/commands/rhoas_kafka"},"rhoas kafka"),"\t - Create, view, use, and manage your Kafka instances"),(0,o.kt)("li",{parentName:"ul"},(0,o.kt)("a",{parentName:"li",href:"/app-services-cli/commands/rhoas_kafka_consumer-group_delete"},"rhoas kafka consumer-group delete"),"\t - Delete a consumer group"),(0,o.kt)("li",{parentName:"ul"},(0,o.kt)("a",{parentName:"li",href:"/app-services-cli/commands/rhoas_kafka_consumer-group_describe"},"rhoas kafka consumer-group describe"),"\t - Describe a consumer group"),(0,o.kt)("li",{parentName:"ul"},(0,o.kt)("a",{parentName:"li",href:"/app-services-cli/commands/rhoas_kafka_consumer-group_list"},"rhoas kafka consumer-group list"),"\t - List all consumer groups"),(0,o.kt)("li",{parentName:"ul"},(0,o.kt)("a",{parentName:"li",href:"/app-services-cli/commands/rhoas_kafka_consumer-group_reset-offset"},"rhoas kafka consumer-group reset-offset"),"\t - Reset partition offsets for a consumer group")))}m.isMDXComponent=!0}}]);