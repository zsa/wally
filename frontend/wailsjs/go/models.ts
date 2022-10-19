export namespace state {
	
	export class Update {
	    version: string;
	    url: string;
	    releaseNotes: string;
	
	    static createFrom(source: any = {}) {
	        return new Update(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.version = source["version"];
	        this.url = source["url"];
	        this.releaseNotes = source["releaseNotes"];
	    }
	}

}

