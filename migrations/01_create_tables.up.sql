CREATE TABLE users (
    id UUID PRIMARY KEY ,
    full_name TEXT DEFAULT '',
    image_url TEXT NOT NULL ,
    phone_number TEXT NOT NULL ,
    email TEXT NOT NULL ,
    occupation TEXT NOT NULL ,
    address TEXT NOT NULL ,
    profile_info TEXT DEFAULT '' ,
);

CREATE TABLE educations (
    id UUID PRIMARY KEY ,
    user_id UUID REFERENCES users(id) ,
    school_name TEXT NOT NULL ,
    speciality TEXT NOT NULL ,
    degree TEXT DEFAULT '' ,
    from DATE DEFAULT CURRENT_DATE ,
    to DATE ,
    description TEXT DEFAULT ''
);

CREATE TABLE experiences (
    id UUID PRIMARY KEY ,
    user_id UUID REFERENCES users(id) ,
    company_name TEXT NOT NULL ,
    role TEXT NOT NULL ,
    from DATE DEFAULT CURRENT_DATE ,
    to DATE ,
    description TEXT DEFAULT ''
);

CREATE TABLE skills (
    id UUID PRIMARY KEY ,
    user_id UUID REFERENCES users(id) ,
    name TEXT NOT NULL ,
    degree TEXT DEFAULT ''
);

CREATE TABLE portfolios (
    id UUID PRIMARY KEY ,
    user_id UUID REFERENCES users(id) ,
    project_name TEXT NOT NULL ,
    link TEXT DEFAULT '' ,
    image_url TEXT DEFAULT '' ,
    description TEXT DEFAULT ''
);

CREATE TABLE links (
    id UUID PRIMARY KEY ,
    user_id UUID REFERENCES users(id) ,
    name TEXT NOT NULL ,
    url TEXT NOT NULL ,
    image_url TEXT NOT NULL
);