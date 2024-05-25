import { Injectable } from '@nestjs/common';
import { InjectModel } from '@nestjs/mongoose';
import { Model } from 'mongoose';
import * as bcrypt from 'bcrypt';
import { Admin } from './schemas/admin.schema';
import { CreateAdminDto } from './dto/create-admin.dto';
import { UpdateAdminDto } from './dto/update-admin.dto';

@Injectable()
export class AdminService {
  constructor(@InjectModel(Admin.name) private adminModel: Model<Admin>) {}

  async create(createAdminDto: CreateAdminDto): Promise<Admin> {
    const hashedPassword = await bcrypt.hash(createAdminDto.password, 10);
    const createdAdmin = new this.adminModel({
      ...createAdminDto,
      password: hashedPassword,
    });
    const result = createdAdmin.save();
    result.then((admin) => {
      admin.password = undefined;
    });
    return result;
  }

  async findAll(): Promise<Admin[]> {
    const result = this.adminModel.find().exec();
    result.then((admins) => {
      admins.forEach((admin) => {
        admin.password = undefined;
      });
    });
    return result;
  }

  async findOne(id: string): Promise<Admin> {
    const result = this.adminModel.findById(id).exec();
    result.then((admin) => {
      admin.password = undefined;
    });
    return result;
  }

  async findByUsername(username: string): Promise<Admin> {
    const result = this.adminModel.findOne({
      username,
    });
    result.then((admin) => {
      admin.password = undefined;
    });
    return result;
  }

  async remove(id: string): Promise<Admin> {
    const result = this.adminModel.findByIdAndDelete(id).exec();
    result.then((admin) => {
      admin.password = undefined;
    });
    return result;
  }

  async update(id: string, updateAdminDto: UpdateAdminDto): Promise<Admin> {
    const hashedPassword = await bcrypt.hash(updateAdminDto.password, 10);
    const result = this.adminModel
      .findByIdAndUpdate(
        id,
        { ...updateAdminDto, password: hashedPassword },
        { new: true },
      )
      .exec();
    result.then((admin) => {
      admin.password = undefined;
    });
    return result;
  }
}
