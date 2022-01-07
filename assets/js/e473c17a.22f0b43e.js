"use strict";(self.webpackChunkwebsite=self.webpackChunkwebsite||[]).push([[674],{3905:function(e,t,r){r.d(t,{Zo:function(){return u},kt:function(){return m}});var n=r(7294);function a(e,t,r){return t in e?Object.defineProperty(e,t,{value:r,enumerable:!0,configurable:!0,writable:!0}):e[t]=r,e}function c(e,t){var r=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);t&&(n=n.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),r.push.apply(r,n)}return r}function o(e){for(var t=1;t<arguments.length;t++){var r=null!=arguments[t]?arguments[t]:{};t%2?c(Object(r),!0).forEach((function(t){a(e,t,r[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(r)):c(Object(r)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(r,t))}))}return e}function i(e,t){if(null==e)return{};var r,n,a=function(e,t){if(null==e)return{};var r,n,a={},c=Object.keys(e);for(n=0;n<c.length;n++)r=c[n],t.indexOf(r)>=0||(a[r]=e[r]);return a}(e,t);if(Object.getOwnPropertySymbols){var c=Object.getOwnPropertySymbols(e);for(n=0;n<c.length;n++)r=c[n],t.indexOf(r)>=0||Object.prototype.propertyIsEnumerable.call(e,r)&&(a[r]=e[r])}return a}var s=n.createContext({}),l=function(e){var t=n.useContext(s),r=t;return e&&(r="function"==typeof e?e(t):o(o({},t),e)),r},u=function(e){var t=l(e.components);return n.createElement(s.Provider,{value:t},e.children)},p={inlineCode:"code",wrapper:function(e){var t=e.children;return n.createElement(n.Fragment,{},t)}},d=n.forwardRef((function(e,t){var r=e.components,a=e.mdxType,c=e.originalType,s=e.parentName,u=i(e,["components","mdxType","originalType","parentName"]),d=l(r),m=a,v=d["".concat(s,".").concat(m)]||d[m]||p[m]||c;return r?n.createElement(v,o(o({ref:t},u),{},{components:r})):n.createElement(v,o({ref:t},u))}));function m(e,t){var r=arguments,a=t&&t.mdxType;if("string"==typeof e||a){var c=r.length,o=new Array(c);o[0]=d;var i={};for(var s in t)hasOwnProperty.call(t,s)&&(i[s]=t[s]);i.originalType=e,i.mdxType="string"==typeof e?e:a,o[1]=i;for(var l=2;l<c;l++)o[l]=r[l];return n.createElement.apply(null,o)}return n.createElement.apply(null,r)}d.displayName="MDXCreateElement"},9561:function(e,t,r){r.r(t),r.d(t,{frontMatter:function(){return i},contentTitle:function(){return s},metadata:function(){return l},toc:function(){return u},default:function(){return d}});var n=r(7462),a=r(3366),c=(r(7294),r(3905)),o=["components"],i={},s=void 0,l={unversionedId:"commands/rhoas_service-account_create",id:"commands/rhoas_service-account_create",isDocsHomePage:!1,title:"rhoas_service-account_create",description:"rhoas service-account create",source:"@site/docs/commands/rhoas_service-account_create.md",sourceDirName:"commands",slug:"/commands/rhoas_service-account_create",permalink:"/app-services-cli/commands/rhoas_service-account_create",editUrl:"https://github.com/redhat-developer/app-services-cli/docs/commands/rhoas_service-account_create.md",tags:[],version:"current",frontMatter:{},sidebar:"tutorialSidebar",previous:{title:"rhoas_service-account",permalink:"/app-services-cli/commands/rhoas_service-account"},next:{title:"rhoas_service-account_delete",permalink:"/app-services-cli/commands/rhoas_service-account_delete"}},u=[{value:"rhoas service-account create",id:"rhoas-service-account-create",children:[{value:"Synopsis",id:"synopsis",children:[]},{value:"Examples",id:"examples",children:[]},{value:"Options",id:"options",children:[]},{value:"Options inherited from parent commands",id:"options-inherited-from-parent-commands",children:[]},{value:"SEE ALSO",id:"see-also",children:[]}]}],p={toc:u};function d(e){var t=e.components,r=(0,a.Z)(e,o);return(0,c.kt)("wrapper",(0,n.Z)({},p,r,{components:t,mdxType:"MDXLayout"}),(0,c.kt)("h2",{id:"rhoas-service-account-create"},"rhoas service-account create"),(0,c.kt)("p",null,"Create a service account"),(0,c.kt)("h3",{id:"synopsis"},"Synopsis"),(0,c.kt)("p",null,"Create a service account with credentials that are saved to a file."),(0,c.kt)("p",null,"Applications and tools use these service account credentials to authenticate and interact with your application services."),(0,c.kt)("p",null,"You must specify an output format into which the credentials will be stored."),(0,c.kt)("ul",null,(0,c.kt)("li",{parentName:"ul"},"env (default): Store credentials in an env file as environment variables"),(0,c.kt)("li",{parentName:"ul"},"json: Store credentials in a JSON file"),(0,c.kt)("li",{parentName:"ul"},"properties: Store credentials in a properties file, which is typically used in Java-related technologies.")),(0,c.kt)("pre",null,(0,c.kt)("code",{parentName:"pre"},"rhoas service-account create [flags]\n")),(0,c.kt)("h3",{id:"examples"},"Examples"),(0,c.kt)("pre",null,(0,c.kt)("code",{parentName:"pre"},"# Create a service account through an interactive prompt\n$ rhoas service-account create\n\n# Create a service account and save the credentials in a JSON file\n$ rhoas service-account create --file-format json\n\n# Create a service account and forcibly overwrite the credentials file if it exists already\n$ rhoas service-account create --overwrite\n\n# Create a service account and save credentials to a custom file location\n$ rhoas service-account create --output-file=./service-acct-credentials.json\n\n")),(0,c.kt)("h3",{id:"options"},"Options"),(0,c.kt)("pre",null,(0,c.kt)("code",{parentName:"pre"},'      --file-format string         Format in which to save the service account credentials (choose from: "env", "json", "properties")\n      --output-file string         Sets a custom file location to save the credentials\n      --overwrite                  Forcibly overwrite a credentials file if it already exists\n      --short-description string   Short description of the service account\n')),(0,c.kt)("h3",{id:"options-inherited-from-parent-commands"},"Options inherited from parent commands"),(0,c.kt)("pre",null,(0,c.kt)("code",{parentName:"pre"},"  -h, --help      Show help for a command\n  -v, --verbose   Enable verbose mode\n")),(0,c.kt)("h3",{id:"see-also"},"SEE ALSO"),(0,c.kt)("ul",null,(0,c.kt)("li",{parentName:"ul"},(0,c.kt)("a",{parentName:"li",href:"/app-services-cli/commands/rhoas_service-account"},"rhoas service-account"),"\t - Create, list, describe, delete, and update service accounts")))}d.isMDXComponent=!0}}]);