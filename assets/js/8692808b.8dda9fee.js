"use strict";(self.webpackChunkwebsite=self.webpackChunkwebsite||[]).push([[9331],{3905:function(e,t,o){o.d(t,{Zo:function(){return f},kt:function(){return m}});var r=o(7294);function n(e,t,o){return t in e?Object.defineProperty(e,t,{value:o,enumerable:!0,configurable:!0,writable:!0}):e[t]=o,e}function s(e,t){var o=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);t&&(r=r.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),o.push.apply(o,r)}return o}function a(e){for(var t=1;t<arguments.length;t++){var o=null!=arguments[t]?arguments[t]:{};t%2?s(Object(o),!0).forEach((function(t){n(e,t,o[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(o)):s(Object(o)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(o,t))}))}return e}function i(e,t){if(null==e)return{};var o,r,n=function(e,t){if(null==e)return{};var o,r,n={},s=Object.keys(e);for(r=0;r<s.length;r++)o=s[r],t.indexOf(o)>=0||(n[o]=e[o]);return n}(e,t);if(Object.getOwnPropertySymbols){var s=Object.getOwnPropertySymbols(e);for(r=0;r<s.length;r++)o=s[r],t.indexOf(o)>=0||Object.prototype.propertyIsEnumerable.call(e,o)&&(n[o]=e[o])}return n}var p=r.createContext({}),c=function(e){var t=r.useContext(p),o=t;return e&&(o="function"==typeof e?e(t):a(a({},t),e)),o},f=function(e){var t=c(e.components);return r.createElement(p.Provider,{value:t},e.children)},u={inlineCode:"code",wrapper:function(e){var t=e.children;return r.createElement(r.Fragment,{},t)}},l=r.forwardRef((function(e,t){var o=e.components,n=e.mdxType,s=e.originalType,p=e.parentName,f=i(e,["components","mdxType","originalType","parentName"]),l=c(o),m=n,d=l["".concat(p,".").concat(m)]||l[m]||u[m]||s;return o?r.createElement(d,a(a({ref:t},f),{},{components:o})):r.createElement(d,a({ref:t},f))}));function m(e,t){var o=arguments,n=t&&t.mdxType;if("string"==typeof e||n){var s=o.length,a=new Array(s);a[0]=l;var i={};for(var p in t)hasOwnProperty.call(t,p)&&(i[p]=t[p]);i.originalType=e,i.mdxType="string"==typeof e?e:n,a[1]=i;for(var c=2;c<s;c++)a[c]=o[c];return r.createElement.apply(null,a)}return r.createElement.apply(null,o)}l.displayName="MDXCreateElement"},7639:function(e,t,o){o.r(t),o.d(t,{frontMatter:function(){return i},contentTitle:function(){return p},metadata:function(){return c},toc:function(){return f},default:function(){return l}});var r=o(7462),n=o(3366),s=(o(7294),o(3905)),a=["components"],i={},p=void 0,c={unversionedId:"commands/rhoas_kafka_consumer-group_reset-offset",id:"commands/rhoas_kafka_consumer-group_reset-offset",isDocsHomePage:!1,title:"rhoas_kafka_consumer-group_reset-offset",description:"rhoas kafka consumer-group reset-offset",source:"@site/docs/commands/rhoas_kafka_consumer-group_reset-offset.md",sourceDirName:"commands",slug:"/commands/rhoas_kafka_consumer-group_reset-offset",permalink:"/app-services-cli/commands/rhoas_kafka_consumer-group_reset-offset",editUrl:"https://github.com/redhat-developer/app-services-cli/docs/commands/rhoas_kafka_consumer-group_reset-offset.md",tags:[],version:"current",frontMatter:{},sidebar:"tutorialSidebar",previous:{title:"rhoas_kafka_consumer-group_list",permalink:"/app-services-cli/commands/rhoas_kafka_consumer-group_list"},next:{title:"rhoas_kafka_create",permalink:"/app-services-cli/commands/rhoas_kafka_create"}},f=[{value:"rhoas kafka consumer-group reset-offset",id:"rhoas-kafka-consumer-group-reset-offset",children:[{value:"Synopsis",id:"synopsis",children:[]},{value:"Examples",id:"examples",children:[]},{value:"Options",id:"options",children:[]},{value:"Options inherited from parent commands",id:"options-inherited-from-parent-commands",children:[]},{value:"SEE ALSO",id:"see-also",children:[]}]}],u={toc:f};function l(e){var t=e.components,o=(0,n.Z)(e,a);return(0,s.kt)("wrapper",(0,r.Z)({},u,o,{components:t,mdxType:"MDXLayout"}),(0,s.kt)("h2",{id:"rhoas-kafka-consumer-group-reset-offset"},"rhoas kafka consumer-group reset-offset"),(0,s.kt)("p",null,"Reset partition offsets for a consumer group"),(0,s.kt)("h3",{id:"synopsis"},"Synopsis"),(0,s.kt)("p",null,"Reset partition offsets for a particular topic. A reset changes the offset position from which consumers read from the message log of a topic partition. To reset an offset position, the consumer group must have NO MEMBERS connected to a topic."),(0,s.kt)("p",null,"You can choose from the following options for the reset:"),(0,s.kt)("ul",null,(0,s.kt)("li",{parentName:"ul"},"Earliest (earliest offset at the start of the message log)"),(0,s.kt)("li",{parentName:"ul"},"Latest (latest offset at the end of the message log)"),(0,s.kt)("li",{parentName:"ul"},"Absolute (specific offset in the message log)"),(0,s.kt)("li",{parentName:"ul"},"Timestamp (specific timestamp in the message log)")),(0,s.kt)("p",null,"You can also reset the offset position for all topics or a single, specified topic."),(0,s.kt)("p",null,"Warning: By resetting the offset position, you risk clients skipping or duplicating messages."),(0,s.kt)("pre",null,(0,s.kt)("code",{parentName:"pre"},"rhoas kafka consumer-group reset-offset [flags]\n")),(0,s.kt)("h3",{id:"examples"},"Examples"),(0,s.kt)("pre",null,(0,s.kt)("code",{parentName:"pre"},'# Reset partition offsets for a consumer group to latest\n$ rhoas kafka consumer-group reset-offset --id consumer_group_1 --topic my-topic --offset latest\n\n# Reset partition offsets for a consumer group to earliest\n$ rhoas kafka consumer-group reset-offset --id consumer_group_1 --topic my-topic --offset earliest\n\n# Reset partition offsets for a consumer group to an absolute value\n$ rhoas kafka consumer-group reset-offset --id consumer_group_1 --topic my-topic --offset absolute --value 0\n\n# Reset partition offsets for a consumer group to a timestamp\n$ rhoas kafka consumer-group reset-offset --id consumer_group_1 --topic my-topic --offset timestamp --value "2016-06-23T09:07:21-07:00"\n\n# Reset specific partition offsets for a consumer group\n$ rhoas kafka consumer-group reset-offset --id consumer_group_1 --topic my-topic --offset latest --partitions 0,1\n\n')),(0,s.kt)("h3",{id:"options"},"Options"),(0,s.kt)("pre",null,(0,s.kt)("code",{parentName:"pre"},'      --id string               The unique ID of the consumer group to reset-offset\n      --offset string           Offset type (choose from: "earliest", "latest", "absolute", "timestamp")\n      --partitions int32Slice   Reset consumer group offsets on specified partitions (comma-separated integers) (default [])\n      --topic string            Reset consumer group offsets on a specified topic\n      --value string            Custom offset value (required when offset is "absolute" or "timestamp")\n  -y, --yes                     Skip confirmation of this action \n')),(0,s.kt)("h3",{id:"options-inherited-from-parent-commands"},"Options inherited from parent commands"),(0,s.kt)("pre",null,(0,s.kt)("code",{parentName:"pre"},"  -h, --help      Show help for a command\n  -v, --verbose   Enable verbose mode\n")),(0,s.kt)("h3",{id:"see-also"},"SEE ALSO"),(0,s.kt)("ul",null,(0,s.kt)("li",{parentName:"ul"},(0,s.kt)("a",{parentName:"li",href:"/app-services-cli/commands/rhoas_kafka_consumer-group"},"rhoas kafka consumer-group"),"\t - Describe, list, and delete consumer groups for the current Kafka instance")))}l.isMDXComponent=!0}}]);